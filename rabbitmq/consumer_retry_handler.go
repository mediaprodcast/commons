package broker

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type ConsumerRetryHandler struct {
	maxRetryCount   int
	exchangeName    string
	dlqExchangeName string
	dlxExchangeName string
	logger          *zap.Logger
}

// Retry handles the retry logic for a message consumed from RabbitMQ.
// It increments the retry count in the message headers and republishes the message.
// If the retry count exceeds the maximum retry count, the message is moved to the Dead Letter Queue (DLQ).
//
// Parameters:
//   - ch: The RabbitMQ channel used for publishing the message.
//   - d: The delivery object containing the message and its metadata.
//
// Returns:
//   - error: An error if the message could not be republished or moved to the DLQ.
func (c *ConsumerRetryHandler) Retry(ch *amqp.Channel, d *amqp.Delivery) error {
	if d.Headers == nil {
		d.Headers = amqp.Table{}
	}

	retryCount, ok := d.Headers["x-retry-count"].(int64)
	if !ok {
		retryCount = 0
	}
	retryCount++
	d.Headers["x-retry-count"] = retryCount

	c.logger.Info("Retrying message", zap.ByteString("body", d.Body), zap.Int64("retryCount", retryCount))

	if retryCount >= int64(c.maxRetryCount) {
		c.logger.Info("Moving message to DLQ", zap.String("DLQ", c.dlqExchangeName))

		return ch.PublishWithContext(context.Background(), "", c.dlqExchangeName, false, false, amqp.Publishing{
			ContentType:  "application/json",
			Headers:      d.Headers,
			Body:         d.Body,
			DeliveryMode: amqp.Persistent,
		})
	}

	time.Sleep(time.Second * time.Duration(retryCount))

	return ch.PublishWithContext(
		context.Background(),
		d.Exchange,
		d.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Headers:      d.Headers,
			Body:         d.Body,
			DeliveryMode: amqp.Persistent,
		},
	)
}

// Declare sets up the main queue, Dead Letter Exchange (DLX), and Dead Letter Queue (DLQ) for the consumer.
// It declares the main queue, the DLX exchange, binds the main queue to the DLX, and declares the DLQ.
//
// Parameters:
//   - ch: The AMQP channel used to declare the queue and exchange.
//
// Returns:
//   - error: An error if any of the declarations or bindings fail.
func (c *ConsumerRetryHandler) Declare(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		c.exchangeName, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare main queue: %w", err)
	}

	// Declare DLX
	err = ch.ExchangeDeclare(
		c.dlxExchangeName, // name
		"fanout",          // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		c.logger.Error("Failed to declare DLX exchange", zap.String("exchange", c.dlxExchangeName), zap.Error(err))
		return fmt.Errorf("failed to declare DLX exchange: %w", err)
	}

	// Bind main queue to DLX (this is incorrect usage; should be in queue arguments)
	err = ch.QueueBind(
		q.Name,            // queue name
		"",                // routing key
		c.dlxExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue to DLX: %w", err)
	}

	// Declare DLQ
	_, err = ch.QueueDeclare(
		c.dlqExchangeName, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare DLQ: %w", err)
	}

	return nil
}
