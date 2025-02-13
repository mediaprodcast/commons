package broker

import (
	"crypto/tls"
	"time"

	"github.com/mediaprodcast/commons/env"
	"go.uber.org/zap"
)

type ConnectionBuilder struct {
	conn *Connection
}

func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{
		conn: &Connection{
			logger:            zap.NewNop(),
			connRetryInterval: 5 * time.Second,
			address:           env.GetString("RABBITMQ_ADDR", "amqp://guest:guest@localhost:5672/"),
		},
	}
}

func (b *ConnectionBuilder) WithConnRetryInterval(interval time.Duration) *ConnectionBuilder {
	b.conn.connRetryInterval = interval
	return b
}

func (b *ConnectionBuilder) WithWithTlsConfig(tlsConfig *tls.Config) *ConnectionBuilder {
	b.conn.tlsConfig = tlsConfig
	return b
}

func (b *ConnectionBuilder) WithAddress(address string) *ConnectionBuilder {
	b.conn.address = address
	return b
}

func (b *ConnectionBuilder) WithLogger(logger *zap.Logger) *ConnectionBuilder {
	b.conn.logger = logger
	return b
}

func (b *ConnectionBuilder) Build() *Connection {
	return b.conn
}
