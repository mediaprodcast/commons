package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	broker "github.com/mediaprodcast/commons/rabbitmq" // Replace with your actual module path
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func messageHandler(ctx context.Context, msg []byte) error {
	fmt.Println(string(msg))

	return nil
}

type Message struct {
	Data string
}

func main() {
	logger, _ := zap.NewProduction()
	tracer := otel.Tracer("amqp")

	duplex := broker.NewDuplexBuilder().
		// Logger
		WithLogger(logger).
		// Tracer
		WithTracer(tracer).
		// Connection
		WithConnection(func(builder *broker.ConnectionBuilder) *broker.ConnectionBuilder {
			return builder.WithAddress("amqp://guest:guest@localhost:5672/")
		}).
		// Consumer
		WithConsumer(func(builder *broker.ConsumerBuilder) *broker.ConsumerBuilder {
			return builder.WithMessageHandler(messageHandler)
		}).
		// Exchange
		WithExchange(func(builder *broker.ExchangeBuilder) *broker.ExchangeBuilder {
			return builder.
				WithExchange(broker.PackagerGeneratePlaylistExchangeName).
				WithKind(broker.DirectExchangeKind)
		}).
		// Queue
		WithQueue(func(builder *broker.QueueBuilder) *broker.QueueBuilder {
			return builder.
				WithName(broker.QueueName(broker.PackagerGeneratePlaylistExchangeName)).
				WithExchange(broker.PackagerGeneratePlaylistExchangeName)
		}).
		// Message retry handler
		WithRetryHandler().
		// Producer
		WithProducer().
		// Duplex
		Build()

	go func() {
		for {
			if duplex.Consumer.Ready() {
				msg := &Message{Data: "Hello world"}

				body, err := json.Marshal(msg)
				if err != nil {
					logger.Error(err.Error())
				}

				// Message
				message := broker.NewMessageBuilder().
					WithExchange(broker.PackagerGeneratePlaylistExchangeName).
					WithKey("").
					WithMandatory(false).
					WithImmediate(false).
					WithDeliveryMode(amqp.Persistent).
					WithContentType("application/json").
					WithBody(body).
					Build()

				err = duplex.Publish(context.Background(), message)
				if err != nil {
					fmt.Println(err.Error())
				}

				time.Sleep(5 * time.Second)
			}
		}
	}()

	duplex.ConnectAndConsume()
}
