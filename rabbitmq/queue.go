package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Name       QueueName // name / event
	Exchange   ExchangeName
	Durable    bool       // durable
	AutoDelete bool       // delete when unused
	Exclusive  bool       // exclusive
	NoWait     bool       // no-wait
	Args       amqp.Table // arguments
	RoutingKey string     // key
}

func (q *Queue) Declare(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		string(q.Name), // name
		q.Durable,      // durable
		q.AutoDelete,   // delete when unused
		q.Exclusive,    // exclusive
		q.NoWait,       // no-wait
		q.Args,         // arguments
	)
}

func (q *Queue) Bind(ch *amqp.Channel) error {
	return ch.QueueBind(
		string(q.Name),     // queue name
		q.RoutingKey,       // routing key
		string(q.Exchange), // exchange
		false,              // no-wait
		nil,
	)
}
