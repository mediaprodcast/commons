package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueBuilder struct {
	queue *Queue
}

func NewQueueBuilder() *QueueBuilder {
	return &QueueBuilder{
		queue: &Queue{},
	}
}

func (b *QueueBuilder) WithName(name QueueName) *QueueBuilder {
	b.queue.Name = name
	return b
}

func (b *QueueBuilder) WithExchange(exchange ExchangeName) *QueueBuilder {
	b.queue.Exchange = exchange
	return b
}

func (b *QueueBuilder) WithDurable(durable bool) *QueueBuilder {
	b.queue.Durable = durable
	return b
}

func (b *QueueBuilder) WithAutoDelete(autoDelete bool) *QueueBuilder {
	b.queue.AutoDelete = autoDelete
	return b
}

func (b *QueueBuilder) WithExclusive(exclusive bool) *QueueBuilder {
	b.queue.Exclusive = exclusive
	return b
}

func (b *QueueBuilder) WithNoWait(noWait bool) *QueueBuilder {
	b.queue.NoWait = noWait
	return b
}

func (b *QueueBuilder) WithArgs(args amqp.Table) *QueueBuilder {
	b.queue.Args = args
	return b
}

func (b *QueueBuilder) WithRoutingKey(routingKey string) *QueueBuilder {
	b.queue.RoutingKey = routingKey
	return b
}

func (b *QueueBuilder) Build() *Queue {
	return b.queue
}
