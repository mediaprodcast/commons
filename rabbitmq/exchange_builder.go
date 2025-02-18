package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type ExchangeBuilder struct {
	exchange *Exchange
}

func NewExchangeBuilder() *ExchangeBuilder {
	return &ExchangeBuilder{
		exchange: &Exchange{},
	}
}

func (b *ExchangeBuilder) WithExchange(name ExchangeName) *ExchangeBuilder {
	b.exchange.Name = name
	return b
}

func (b *ExchangeBuilder) WithKind(kind ExchangeKind) *ExchangeBuilder {
	b.exchange.Kind = kind
	return b
}

func (b *ExchangeBuilder) WithDurable(durable bool) *ExchangeBuilder {
	b.exchange.Durable = durable
	return b
}

func (b *ExchangeBuilder) WithAutoDelete(autoDelete bool) *ExchangeBuilder {
	b.exchange.AutoDelete = autoDelete
	return b
}

func (b *ExchangeBuilder) WithExclusive(exclusive bool) *ExchangeBuilder {
	b.exchange.Exclusive = exclusive
	return b
}

func (b *ExchangeBuilder) WithNoWait(noWait bool) *ExchangeBuilder {
	b.exchange.NoWait = noWait
	return b
}

func (b *ExchangeBuilder) WithInternal(internal bool) *ExchangeBuilder {
	b.exchange.Internal = internal
	return b
}

func (b *ExchangeBuilder) WithArgs(args amqp.Table) *ExchangeBuilder {
	b.exchange.Args = args
	return b
}

func (b *ExchangeBuilder) Build() *Exchange {
	return b.exchange
}
