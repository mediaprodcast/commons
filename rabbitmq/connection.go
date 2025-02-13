package broker

import (
	"crypto/tls"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Connection struct {
	connRetryInterval time.Duration
	conn              *amqp.Connection
	ch                *amqp.Channel
	address           string
	tlsConfig         *tls.Config
	logger            *zap.Logger

	closeOnce sync.Once
	done      chan struct{} // Channel to stop watching
	mu        sync.Mutex    // Ensures thread safety
}

func (c *Connection) Connect() (*amqp.Connection, error) {
	if c.Ready() {
		return c.conn, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.cleanupLocked()

	return c.connect()
}

func (c *Connection) ConnectWithReconnect() error {
	if c.Ready() {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.cleanupLocked()

	for {
		if _, err := c.connect(); err == nil {
			c.done = make(chan struct{})
			go c.watchConnection()
			return nil
		}
		c.logger.Warn("Retrying connection to RabbitMQ...", zap.Duration("retry_in", c.connRetryInterval))
		time.Sleep(c.connRetryInterval)
	}
}

func (c *Connection) connect() (*amqp.Connection, error) {
	conn, err := c.dial(c.address)
	if err != nil {
		c.logger.Error("Failed to connect to RabbitMQ", zap.Error(err))
		return nil, err
	}
	c.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		c.logger.Error("Failed to open a channel", zap.Error(err))
		conn.Close()
		return nil, err
	}
	c.ch = ch
	c.logger.Info("Connected to RabbitMQ successfully")
	return conn, nil
}

func (c *Connection) dial(address string) (*amqp.Connection, error) {
	if c.tlsConfig != nil {
		return amqp.DialTLS(address, c.tlsConfig)
	}
	return amqp.Dial(address)
}

func (c *Connection) watchConnection() {
	closeChan := make(chan *amqp.Error, 1)
	c.conn.NotifyClose(closeChan)

	select {
	case err := <-closeChan:
		if err != nil {
			c.logger.Warn("Connection lost. Reconnecting...", zap.Error(err))
			c.Reconnect()
		}
	case <-c.done:
	}
}

func (c *Connection) Reconnect() {
	c.Cleanup()
	time.Sleep(c.connRetryInterval)
	c.Connect()
}

func (c *Connection) Ready() bool {
	return c.IsConnected() && c.IsChannelOpen()
}

func (c *Connection) IsConnected() bool {
	return c.conn != nil && !c.conn.IsClosed()
}

func (c *Connection) IsChannelOpen() bool {
	return c.ch != nil && !c.ch.IsClosed()
}

func (c *Connection) Cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cleanupLocked()
}

func (c *Connection) cleanupLocked() {
	c.closeOnce.Do(func() {
		if c.done != nil {
			close(c.done)
			c.done = nil
		}
		if c.ch != nil {
			c.ch.Close()
			c.ch = nil
		}
		if c.conn != nil {
			c.conn.Close()
			c.conn = nil
		}
	})
}
