package broker

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type DuplexBuilder struct {
	connectionBuilder   *ConnectionBuilder
	consumerBuilder     *ConsumerBuilder
	producerBuilder     *ProducerBuilder
	exchangeBuilder     *ExchangeBuilder
	queueBuilder        *QueueBuilder
	logger              *zap.Logger
	tracer              trace.Tracer
	retryHandlerBuilder *ConsumerRetryHandlerBuilder
}

func NewDuplexBuilder() *DuplexBuilder {
	return &DuplexBuilder{
		logger:            zap.L(),
		connectionBuilder: NewConnectionBuilder(),
		producerBuilder:   NewProducerBuilder(nil),
		tracer:            otel.Tracer("amqp"),
	}
}

// WithConnection allows chaining connection configuration
func (d *DuplexBuilder) WithConnection(fn func(*ConnectionBuilder) *ConnectionBuilder) *DuplexBuilder {
	d.connectionBuilder = fn(d.connectionBuilder)

	return d
}

// WithConsumer allows chaining consumer configuration
func (d *DuplexBuilder) WithConsumer(fn func(*ConsumerBuilder) *ConsumerBuilder) *DuplexBuilder {
	if d.consumerBuilder == nil {
		d.consumerBuilder = NewConsumerBuilder(nil)
	}

	d.consumerBuilder = fn(d.consumerBuilder)
	return d
}

// WithProducer allows chaining producer configuration
func (d *DuplexBuilder) WithProducer(fns ...func(*ProducerBuilder) *ProducerBuilder) *DuplexBuilder {
	for _, fn := range fns {
		d.producerBuilder = fn(d.producerBuilder)
	}

	return d
}

// WithExchange allows chaining exchange configuration.  Provides defaults if not configured.
func (d *DuplexBuilder) WithExchange(fn func(*ExchangeBuilder) *ExchangeBuilder) *DuplexBuilder {
	if d.exchangeBuilder == nil {
		d.exchangeBuilder = NewExchangeBuilder().
			WithDurable(true).
			WithAutoDelete(false).
			WithInternal(false).
			WithNoWait(false).
			WithArgs(nil)
	}

	d.exchangeBuilder = fn(d.exchangeBuilder)
	return d
}

// WithQueue allows chaining queue configuration. Provides defaults if not configured.
func (d *DuplexBuilder) WithQueue(fn func(*QueueBuilder) *QueueBuilder) *DuplexBuilder {
	if d.queueBuilder == nil {
		d.queueBuilder = NewQueueBuilder().
			WithDurable(true).
			WithAutoDelete(false).
			WithNoWait(false).
			WithArgs(nil)
	}

	d.queueBuilder = fn(d.queueBuilder)
	return d
}

// WithRetryHandler allows chaining retry handler configuration. Provides defaults if not configured.
func (d *DuplexBuilder) WithRetryHandler(fns ...func(*ConsumerRetryHandlerBuilder) *ConsumerRetryHandlerBuilder) *DuplexBuilder {
	if d.retryHandlerBuilder == nil {
		d.retryHandlerBuilder = NewConsumerRetryHandlerBuilder().
			WithMaxRetryCount(100)
	}

	for _, fn := range fns {
		d.retryHandlerBuilder = fn(d.retryHandlerBuilder)
	}

	return d
}

// WithLogger sets the logger.
func (d *DuplexBuilder) WithLogger(logger *zap.Logger) *DuplexBuilder {
	d.logger = logger
	return d
}

// WithTracer sets the tracer.
func (d *DuplexBuilder) WithTracer(tracer trace.Tracer) *DuplexBuilder {
	d.tracer = tracer
	return d
}

// Build creates the Duplex instance.
func (d *DuplexBuilder) Build() *Duplex {
	connection := d.connectionBuilder.
		WithLogger(d.logger).
		Build()

	producer := d.producerBuilder.
		WithLogger(d.logger).
		WithTracer(d.tracer).
		Build()

	producer.conn = connection

	duplex := &Duplex{Producer: producer}

	if d.consumerBuilder != nil {
		consumerBuilder := d.consumerBuilder.
			WithLogger(d.logger).
			WithTracer(d.tracer)

		if d.retryHandlerBuilder != nil {
			consumerBuilder.WithRetryHandler(d.retryHandlerBuilder.WithLogger(d.logger).Build())
		}

		if d.exchangeBuilder != nil {
			consumerBuilder.WithExchange(d.exchangeBuilder.Build())
		}

		if d.queueBuilder != nil {
			consumerBuilder.WithQueue(d.queueBuilder.Build())
		}

		consumer := consumerBuilder.Build()
		consumer.conn = connection
		duplex.Consumer = consumer
	}

	return duplex
}
