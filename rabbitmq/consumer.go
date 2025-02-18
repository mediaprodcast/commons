package broker

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type MessageHandler = func(context.Context, []byte) error

type Consumer struct {
	conn            *Connection
	exchange        *Exchange
	queue           *Queue
	tracer          trace.Tracer
	retryHandler    *ConsumerRetryHandler
	messageHandlers []MessageHandler
	logger          *zap.Logger
	stopChan        chan struct{} // Channel for stopping the consumer
	ready           atomic.Bool
	concurrency     int
}

// ConnectAndConsume keeps the consumer running indefinitely with auto-reconnect
func (c *Consumer) ConnectAndConsume() {
	defer c.conn.Cleanup()

	for {
		// Attempt to connect
		conn, err := c.conn.Connect()
		if err != nil {
			c.logger.Error("Failed to connect to RabbitMQ, retrying...", zap.Error(err))
			time.Sleep(c.conn.connRetryInterval)
			continue // Retry connection
		}

		// Create a channel for connection close events
		closeChan := make(chan *amqp.Error, 1)
		conn.NotifyClose(closeChan)

		// Start consuming messages
		go func() {
			if err := c.consumeOnce(); err != nil {
				c.logger.Error("Error during message consumption", zap.Error(err))
			}
		}()

		// Block until the connection is lost
		err = <-closeChan
		c.logger.Warn("Connection to RabbitMQ lost. Reconnecting...", zap.Error(err))
		c.ready.Store(false) // Connection lost, set ready to false

		// Wait before retrying to avoid rapid reconnection attempts
		time.Sleep(c.conn.connRetryInterval)
	}
}

func (c *Consumer) Ready() bool {
	return c.conn.Ready() && c.ready.Load()
}

// consumeOnce sets up the queue and starts consuming messages
func (c *Consumer) consumeOnce() error {
	if c.exchange != nil {
		if err := c.exchange.Declare(c.conn.ch); err != nil {
			return fmt.Errorf("failed to declare exchange: %w", err)
		}

		c.logger.Info(
			"declared exchange",
			zap.String("exchange", string(c.exchange.Name)),
			zap.String("kind", string(c.exchange.Kind)),
		)
	}

	if c.retryHandler != nil {
		if err := c.retryHandler.Declare(c.conn.ch); err != nil {
			return fmt.Errorf("failed to declare retry handler: %w", err)
		}
	}

	if c.queue != nil {
		q, err := c.queue.Declare(c.conn.ch)
		if err != nil {
			return fmt.Errorf("failed to declare queue: %w", err)
		}

		c.logger.Info(
			"declared queue",
			zap.String("name", string(c.queue.Name)),
			zap.String("key", string(c.queue.RoutingKey)),
		)

		if err := c.queue.Bind(c.conn.ch); err != nil {
			return fmt.Errorf("failed to bind queue: %w", err)
		}

		messages, err := c.conn.ch.Consume(q.Name, "", false, false, false, false, nil)
		if err != nil {
			c.logger.Error("Failed to consume queue", zap.String("queue", q.Name), zap.Error(err))
			return fmt.Errorf("failed to consume: %w", err)
		}

		c.ready.Store(true) // Set ready to true after successful consumption setup

		// Process messages
		for message := range messages {
			c.processMessage(message, q.Name)
		}
	}

	return nil
}

// processMessage handles individual message processing with tracing and error handling
func (c *Consumer) processMessage(message amqp.Delivery, queueName string) {
	c.logger.Info("Received message", zap.String("queue", queueName), zap.Any("body", string(message.Body)))

	ctx := ExtractAMQPHeader(context.Background(), message.Headers)
	_, messageSpan := c.tracer.Start(ctx, fmt.Sprintf("AMQP - consume - %s", queueName))
	defer messageSpan.End()

	var handlerErrors []error

	// Handle messages
	for _, handler := range c.messageHandlers {
		if err := handler(ctx, message.Body); err != nil {
			c.logger.Error("Error handling message", zap.Error(err))
			messageSpan.RecordError(err)
			handlerErrors = append(handlerErrors, err)
		}
	}

	// Handle retry
	if len(handlerErrors) > 0 {
		if c.retryHandler != nil {
			if err := c.retryHandler.Retry(c.conn.ch, &message); err != nil {
				log.Printf("Error handling retry: %v", err)
			}
		}
		return
	}

	message.Ack(false)
	c.logger.Info("Message successfully processed", zap.String("queue", queueName))
}

// Stop gracefully stops the consumer
func (c *Consumer) Stop() {
	close(c.stopChan)
	c.ready.Store(false)
}
