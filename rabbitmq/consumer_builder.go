package broker

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
)

type ConsumerBuilder struct {
	consumer *Consumer
}

func NewConsumerBuilder(conn *Connection) *ConsumerBuilder {
	return &ConsumerBuilder{
		consumer: &Consumer{
			conn:            conn,
			logger:          zap.NewNop(),
			stopChan:        make(chan struct{}),
			tracer:          noop.NewTracerProvider().Tracer("amqp"),
			retryHandler:    NewConsumerRetryHandlerBuilder().Build(),
			messageHandlers: make([]func(context.Context, []byte) error, 0),
		},
	}
}

func (b *ConsumerBuilder) WithRetryHandler(handler *ConsumerRetryHandler) *ConsumerBuilder {
	b.consumer.retryHandler = handler
	return b
}

func (b *ConsumerBuilder) WithTracer(tracer trace.Tracer) *ConsumerBuilder {
	b.consumer.tracer = tracer
	return b
}

func (b *ConsumerBuilder) WithLogger(logger *zap.Logger) *ConsumerBuilder {
	b.consumer.logger = logger
	return b
}

func (b *ConsumerBuilder) WithExchange(exchange *Exchange) *ConsumerBuilder {
	b.consumer.exchange = exchange
	return b
}

func (b *ConsumerBuilder) WithQueue(queue *Queue) *ConsumerBuilder {
	b.consumer.queue = queue
	return b
}

func (b *ConsumerBuilder) WithMessageHandler(messageHandlers ...MessageHandler) *ConsumerBuilder {
	b.consumer.messageHandlers = append(b.consumer.messageHandlers, messageHandlers...)
	return b
}

func (b *ConsumerBuilder) Build() *Consumer {
	// Validation: Ensure required fields are set
	if b.consumer.exchange == nil {
		panic("exchange is required") // Or handle this differently
	}
	if b.consumer.queue == nil {
		panic("queue is required") // Or handle this differently
	}

	return b.consumer
}
