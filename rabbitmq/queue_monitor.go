package broker

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

// QueueMetric contains key metrics for scaling decisions
type QueueMetric struct {
	MessageCount  int // Current messages awaiting processing
	ConsumerCount int // Active consumers on the queue
}

type QueueMetrics = map[QueueName]QueueMetric

type QueueMetricsHandler = func(*QueueMetrics) error

// QueueMonitor manages monitoring of RabbitMQ queues
type QueueMonitor struct {
	conn            *Connection
	monitored       []*Queue
	handlers        []QueueMetricsHandler
	pollingInterval time.Duration
	metrics         QueueMetrics
	stopChan        chan struct{}
	mu              sync.RWMutex
	logger          *zap.Logger
}

func (qm *QueueMonitor) Listen() {
	if !qm.conn.Ready() {
		qm.conn.Connect()
	}

	ticker := time.NewTicker(qm.pollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			qm.updateMetrics()
		case <-qm.stopChan:
			return
		}
	}
}

func (qm *QueueMonitor) Stop() {
	close(qm.stopChan)
}

func (qm *QueueMonitor) GetMetrics() map[QueueName]QueueMetric {
	qm.mu.RLock()
	defer qm.mu.RUnlock()

	return qm.metrics
}

// Updated metrics collection using QueueDeclare
func (qm *QueueMonitor) updateMetrics() {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	for _, item := range qm.monitored {
		if !qm.conn.Ready() {
			qm.logger.Error("Connection is not ready")
			continue
		}

		stats, err := item.Declare(qm.conn.ch)
		if err != nil {
			qm.logger.Error(
				"failed to declare queue",
				zap.String("name", string(item.Name)),
				zap.Error(err),
			)

			continue
		}

		qm.metrics[item.Name] = QueueMetric{
			MessageCount:  stats.Messages,
			ConsumerCount: stats.Consumers,
		}

		// sent stats to handlers
		for _, handler := range qm.handlers {
			if err := handler(&qm.metrics); err != nil {
				qm.logger.Error("Error handling queue metrics", zap.Error(err))
			}
		}
	}
}
