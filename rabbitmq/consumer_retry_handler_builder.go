package broker

import "go.uber.org/zap"

type ConsumerRetryHandlerBuilder struct {
	handler *ConsumerRetryHandler
}

func NewConsumerRetryHandlerBuilder() *ConsumerRetryHandlerBuilder {
	return &ConsumerRetryHandlerBuilder{
		handler: &ConsumerRetryHandler{
			maxRetryCount:   10,
			exchangeName:    "main_queue",
			dlqExchangeName: "dlq_main",
			dlxExchangeName: "dlx_main",
			logger:          zap.NewNop(),
		},
	}
}

func (b *ConsumerRetryHandlerBuilder) WithLogger(logger *zap.Logger) *ConsumerRetryHandlerBuilder {
	b.handler.logger = logger
	return b
}

func (b *ConsumerRetryHandlerBuilder) WithMaxRetryCount(count int) *ConsumerRetryHandlerBuilder {
	b.handler.maxRetryCount = count
	return b
}

func (b *ConsumerRetryHandlerBuilder) WithExchangeName(name string) *ConsumerRetryHandlerBuilder {
	b.handler.exchangeName = name
	return b
}

func (b *ConsumerRetryHandlerBuilder) WithDLQExchangeName(name string) *ConsumerRetryHandlerBuilder {
	b.handler.dlqExchangeName = name
	return b
}

func (b *ConsumerRetryHandlerBuilder) WithDLXExchangeName(name string) *ConsumerRetryHandlerBuilder {
	b.handler.dlxExchangeName = name
	return b
}

func (b *ConsumerRetryHandlerBuilder) Build() *ConsumerRetryHandler {
	return b.handler
}
