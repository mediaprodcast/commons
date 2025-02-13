package broker

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
)

type SampleMessage struct {
	Data string
}

func TestRabbitMQIntegration(t *testing.T) {
	amqpURL := setupRabbitMQ(t)

	// Connection
	connection := NewConnectionBuilder().
		WithAddress(amqpURL).
		Build()
	assert.NotNil(t, connection)

	_, err := connection.Connect()
	assert.NoError(t, err)

	// Exchange
	exchange := NewExchangeBuilder().
		WithExchange(PackagerGeneratePlaylistExchangeName).
		WithKind(DirectExchangeKind).
		WithDurable(true).
		WithAutoDelete(false).
		WithInternal(false).
		WithNoWait(false).
		WithArgs(nil).
		Build()
	assert.NotNil(t, exchange)

	// Queue
	queue := NewQueueBuilder().
		WithName(QueueName(PackagerGeneratePlaylistExchangeName)).
		WithExchange(PackagerGeneratePlaylistExchangeName).
		WithDurable(true).
		WithAutoDelete(false).
		WithNoWait(false).
		WithArgs(nil).
		Build()
	assert.NotNil(t, queue)

	// Setup synchronization
	var receivedMsg []byte
	done := make(chan struct{})

	// Consumer
	handler := func(ctx context.Context, msg []byte) error {
		t.Logf("Received message: %s", string(msg))
		receivedMsg = msg
		close(done)
		return nil
	}

	consumer := NewConsumerBuilder(connection).
		WithExchange(exchange).
		WithQueue(queue).
		WithMessageHandler(handler).
		Build()
	assert.NotNil(t, consumer)

	// Start consumer
	go func() {
		consumer.ConnectAndConsume()
	}()

	// Wait for consumer to be ready with timeout
	timeout := time.After(10 * time.Second)
	ready := false

	for !ready {
		select {
		case <-timeout:
			t.Fatal("Consumer readiness check timed out")
		case <-time.After(1 * time.Second):
			if consumer.Ready() {
				ready = true
			}
		}
	}

	// Producer (same as before)
	producer := NewProducerBuilder(connection).
		WithTracer(otel.Tracer("amqp")).
		Build()
	assert.NotNil(t, producer)

	testBody := &SampleMessage{Data: "Test Message"}
	body, err := json.Marshal(testBody)
	assert.NoError(t, err)

	message := NewMessageBuilder().
		WithExchange(PackagerGeneratePlaylistExchangeName).
		WithKey("").
		WithMandatory(false).
		WithImmediate(false).
		WithDeliveryMode(amqp.Persistent).
		WithContentType("application/json").
		WithBody(body).
		Build()

	// Publish message
	err = producer.Publish(context.Background(), message)
	assert.NoError(t, err)

	// Wait for message or timeout
	select {
	case <-done:
	case <-time.After(20 * time.Second):
		t.Fatal("Timed out waiting for message")
	}

	// Verify message content
	var parsed SampleMessage
	err = json.Unmarshal(receivedMsg, &parsed)
	assert.NoError(t, err)

	assert.Equal(t, testBody.Data, parsed.Data)
}
