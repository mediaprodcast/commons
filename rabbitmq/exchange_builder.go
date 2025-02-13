package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type ExchangeBuilder struct {
	options *Exchange
}

func NewExchangeBuilder() *ExchangeBuilder {
	return &ExchangeBuilder{
		options: &Exchange{},
	}
}

func (b *ExchangeBuilder) WithExchange(name ExchangeName) *ExchangeBuilder {
	b.options.Name = name
	return b
}

func (b *ExchangeBuilder) WithKind(kind ExchangeKind) *ExchangeBuilder {
	b.options.Kind = kind
	return b
}

func (b *ExchangeBuilder) WithDurable(durable bool) *ExchangeBuilder {
	b.options.Durable = durable
	return b
}

func (b *ExchangeBuilder) WithAutoDelete(autoDelete bool) *ExchangeBuilder {
	b.options.AutoDelete = autoDelete
	return b
}

func (b *ExchangeBuilder) WithExclusive(exclusive bool) *ExchangeBuilder {
	b.options.Exclusive = exclusive
	return b
}

func (b *ExchangeBuilder) WithNoWait(noWait bool) *ExchangeBuilder {
	b.options.NoWait = noWait
	return b
}

func (b *ExchangeBuilder) WithInternal(internal bool) *ExchangeBuilder {
	b.options.Internal = internal
	return b
}

func (b *ExchangeBuilder) WithArgs(args amqp.Table) *ExchangeBuilder {
	b.options.Args = args
	return b
}

func (b *ExchangeBuilder) Build() *Exchange {
	return b.options
}
