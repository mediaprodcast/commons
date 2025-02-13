package main

import (
	"fmt"
	"time"

	broker "github.com/mediaprodcast/commons/rabbitmq" // Replace with your actual module path
	"go.uber.org/zap"
)

func monitorHandler(metrics *broker.QueueMetrics) error {
	fmt.Println(metrics)

	return nil
}

func main() {
	logger, _ := zap.NewProduction()

	// Connection
	connection := broker.NewConnectionBuilder().
		WithAddress("amqp://guest:guest@localhost:5672/").
		WithLogger(logger).
		Build()

	// Queue
	queue := broker.NewQueueBuilder().
		WithName(broker.QueueName(broker.PackagerGeneratePlaylistExchangeName)).
		WithDurable(true).
		WithAutoDelete(false).
		WithNoWait(false).
		WithArgs(nil).
		Build()

	monitor := broker.NewQueueMonitorBuilder(connection).
		WithLogger(logger).
		WithQueue(queue).
		WithPollingInterval(4 * time.Second).
		WithQueueMetricsHandler(monitorHandler).
		Build()

	monitor.Listen()
}
