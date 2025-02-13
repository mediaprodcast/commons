package broker

import (
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/zap"
)

type ProducerBuilder struct {
	producer *Producer
}

func NewProducerBuilder(conn *Connection) *ProducerBuilder {
	return &ProducerBuilder{
		producer: &Producer{
			conn:   conn,
			logger: zap.NewNop(),
			tracer: noop.NewTracerProvider().Tracer("amqp"),
		},
	}
}

func (b *ProducerBuilder) WithLogger(logger *zap.Logger) *ProducerBuilder {
	b.producer.logger = logger
	return b
}

func (b *ProducerBuilder) WithTracer(tracer trace.Tracer) *ProducerBuilder {
	b.producer.tracer = tracer
	return b
}

func (b *ProducerBuilder) Build() *Producer {
	return b.producer
}
