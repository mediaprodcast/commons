package broker

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/rabbitmq"
)

func setupRabbitMQ(t *testing.T) string {
	t.Helper()
	ctx := context.Background()

	// Start a RabbitMQ container
	rmqContainer, err := rabbitmq.Run(ctx, "rabbitmq:3.12.11-management-alpine",
		rabbitmq.WithAdminPassword("guest"),
		rabbitmq.WithAdminUsername("guest"),
	)

	if err != nil {
		t.Fatalf("Could not start RabbitMQ container: %s", err)
	}

	// Get the AMQP connection string
	amqpURL, err := rmqContainer.AmqpURL(ctx)
	if err != nil {
		t.Fatalf("Could not get AMQP URL: %s", err)
	}

	// Cleanup function to terminate the container after test

	t.Cleanup(func() {
		if err := rmqContainer.Terminate(ctx); err != nil {
			t.Fatalf("could not terminate container: %v", err)
		}
	})

	return amqpURL
}

// func TestRabbitMQPublishConsume(t *testing.T) {
// 	_, amqpURL := setupRabbitMQ(t)

// 	conn, err := amqp.Dial(amqpURL)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to RabbitMQ: %s", err)
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		t.Fatalf("Failed to open a channel: %s", err)
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"test_queue",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatalf("Failed to declare queue: %s", err)
// 	}

// 	testMsg := "Hello, RabbitMQ!"

// 	err = ch.PublishWithContext(
// 		context.Background(),
// 		"",
// 		q.Name,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(testMsg),
// 		},
// 	)
// 	if err != nil {
// 		t.Fatalf("Failed to publish message: %s", err)
// 	}

// 	// Give RabbitMQ some time to process
// 	time.Sleep(1 * time.Second)

// 	msgs, err := ch.Consume(
// 		q.Name,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatalf("Failed to consume messages: %s", err)
// 	}

// 	select {
// 	case msg := <-msgs:
// 		assert.Equal(t, testMsg, string(msg.Body))
// 	case <-time.After(5 * time.Second):
// 		t.Fatal("Did not receive message in time")
// 	}
// }
