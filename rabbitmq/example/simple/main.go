package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	broker "github.com/mediaprodcast/commons/rabbitmq" // Replace with your actual module path
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func messageHandler(ctx context.Context, msg []byte) error {
	// fmt.Println(string(msg))
	// return nil

	return errors.New("hello world")
}

type Message struct {
	Data string
}

func main() {
	logger, _ := zap.NewProduction()

	// Connection
	connection := broker.NewConnectionBuilder().
		WithAddress("amqp://guest:guest@localhost:5672/").
		WithLogger(logger).
		Build()

	// Exchange
	exchange := broker.NewExchangeBuilder().
		WithExchange(broker.PackagerGeneratePlaylistExchangeName).
		WithKind(broker.DirectExchangeKind).
		WithDurable(true).
		WithAutoDelete(false).
		WithInternal(false).
		WithNoWait(false).
		WithArgs(nil).
		Build()

	// Queue
	queue := broker.NewQueueBuilder().
		WithName(broker.QueueName(broker.PackagerGeneratePlaylistExchangeName)).
		WithExchange(broker.PackagerGeneratePlaylistExchangeName).
		WithDurable(true).
		WithAutoDelete(false).
		WithNoWait(false).
		WithArgs(nil).
		Build()

	// Message retry handler
	retryHandler := broker.NewConsumerRetryHandlerBuilder().
		WithLogger(logger).
		WithMaxRetryCount(2).
		Build()

	// Consumer
	consumer := broker.NewConsumerBuilder(connection).
		WithTracer(otel.Tracer("amqp")).
		WithRetryHandler(retryHandler).
		WithMessageHandler(messageHandler).
		WithExchange(exchange).
		WithLogger(logger).
		WithQueue(queue).
		Build()

	producer := broker.NewProducerBuilder(connection).
		WithTracer(otel.Tracer("amqp")).
		Build()

	go func() {
		for {
			if consumer.Ready() {
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

				err = producer.Publish(context.Background(), message)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				fmt.Println("Sent message")
				break
			}
		}
	}()

	consumer.ConnectAndConsume()
}
