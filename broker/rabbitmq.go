package broker

import (
	"context"
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mediaprodcast/commons/env"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const MaxRetryCount = 3
const DLQ = "dlq_main"

var (
	amqpUser = env.GetString("RABBITMQ_USER", "guest")     // RabbitMQ username
	amqpPass = env.GetString("RABBITMQ_PASS", "guest")     // RabbitMQ password
	amqpHost = env.GetString("RABBITMQ_HOST", "localhost") // RabbitMQ host address
	amqpPort = env.GetString("RABBITMQ_PORT", "5672")      // RabbitMQ port
)

func Connect() (*amqp.Channel, func() error) {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPass, amqpHost, amqpPort)

	conn, err := amqp.Dial(address)
	if err != nil {
		zap.L().Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}

	ch, err := conn.Channel()
	if err != nil {
		zap.L().Fatal("Failed to open a channel", zap.Error(err))
	}

	err = ch.ExchangeDeclare(EventPublishedCount, "direct", true, false, false, false, nil)
	if err != nil {
		zap.L().Fatal("Failed to declare exchange", zap.String("exchange", EventPublishedCount), zap.Error(err))
	}

	err = ch.ExchangeDeclare(EventProcessingErrors, "fanout", true, false, false, false, nil)
	if err != nil {
		zap.L().Fatal("Failed to declare exchange", zap.String("exchange", EventProcessingErrors), zap.Error(err))
	}

	err = createDLQAndDLX(ch)
	if err != nil {
		zap.L().Fatal("Failed to create DLQ and DLX", zap.Error(err))
	}

	return ch, conn.Close
}

func HandleRetry(ch *amqp.Channel, d *amqp.Delivery) error {
	if d.Headers == nil {
		d.Headers = amqp.Table{}
	}

	retryCount, ok := d.Headers["x-retry-count"].(int64)
	if !ok {
		retryCount = 0
	}
	retryCount++
	d.Headers["x-retry-count"] = retryCount

	zap.L().Info("Retrying message", zap.ByteString("body", d.Body), zap.Int64("retryCount", retryCount))

	if retryCount >= MaxRetryCount {
		zap.L().Info("Moving message to DLQ", zap.String("DLQ", DLQ))

		return ch.PublishWithContext(context.Background(), "", DLQ, false, false, amqp.Publishing{
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

func createDLQAndDLX(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		"main_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		zap.L().Fatal("Failed to declare queue", zap.String("queue", "main_queue"), zap.Error(err))
		return err
	}

	// Declare DLX
	dlx := "dlx_main"
	err = ch.ExchangeDeclare(
		dlx,      // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		zap.L().Fatal("Failed to declare DLX exchange", zap.String("exchange", dlx), zap.Error(err))
		return err
	}

	// Bind main queue to DLX
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		dlx,    // exchange
		false,
		nil,
	)
	if err != nil {
		zap.L().Fatal("Failed to bind queue to DLX", zap.String("queue", q.Name), zap.String("exchange", dlx), zap.Error(err))
		return err
	}

	// Declare DLQ
	_, err = ch.QueueDeclare(
		DLQ,   // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		zap.L().Fatal("Failed to declare DLQ", zap.String("queue", DLQ), zap.Error(err))
		return err
	}

	return err
}

type AmqpHeaderCarrier map[string]interface{}

func (a AmqpHeaderCarrier) Get(k string) string {
	value, ok := a[k]
	if !ok {
		return ""
	}

	return value.(string)
}

func (a AmqpHeaderCarrier) Set(k string, v string) {
	a[k] = v
}

func (a AmqpHeaderCarrier) Keys() []string {
	keys := make([]string, len(a))
	i := 0

	for k := range a {
		keys[i] = k
		i++
	}

	return keys
}

func InjectAMQPHeaders(ctx context.Context) map[string]interface{} {
	carrier := make(AmqpHeaderCarrier)
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	return carrier
}

func ExtractAMQPHeader(ctx context.Context, headers map[string]interface{}) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, AmqpHeaderCarrier(headers))
}
