package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Exchange struct {
	Name       ExchangeName // name / event
	Kind       ExchangeKind // kind
	Durable    bool         // durable
	AutoDelete bool         // delete when unused
	Exclusive  bool         // exclusive
	NoWait     bool         // no-wait
	Internal   bool         // internal
	Args       amqp.Table   // arguments
}

func (e *Exchange) Declare(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		string(e.Name),
		string(e.Kind),
		e.Durable,
		e.AutoDelete,
		e.Internal,
		e.NoWait,
		e.Args,
	)
}
