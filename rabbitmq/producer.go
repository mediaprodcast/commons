package broker

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Producer struct {
	conn   *Connection
	logger *zap.Logger
	tracer trace.Tracer
}

func (p *Producer) Publish(ctx context.Context, cfg *Message) error {
	// Try connecting if not connected
	if !p.conn.Ready() {
		_, err := p.conn.Connect()
		if err != nil {
			return err
		}
	}

	ctx, messageSpan := p.tracer.Start(ctx, fmt.Sprintf("AMQP - publish - %s", cfg.Exchange))
	defer messageSpan.End()

	// if queue is not defined, use exchange
	if cfg.Queue != "" {
		cfg.Queue = QueueName(cfg.Exchange)
	}

	// Declare the queue if it doesn't exist
	_, err := p.conn.ch.QueueDeclare(
		string(cfg.Queue), // queue name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		p.logger.Error("Failed to declare queue", zap.String("queue", string(cfg.Queue)), zap.Error(err))
		return err
	}

	if cfg.Publishing.Headers != nil {
		cfg.Publishing.Headers = InjectAMQPHeaders(ctx)
	}

	return p.conn.ch.PublishWithContext(
		ctx,
		string(cfg.Exchange),
		cfg.Key,
		cfg.Mandatory, // mandatory
		cfg.Immediate, // immediate
		*cfg.Publishing,
	)
}
