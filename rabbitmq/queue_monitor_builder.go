package broker

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

type QueueMonitorBuilder struct {
	monitor *QueueMonitor
}

func NewQueueMonitorBuilder(conn *Connection) *QueueMonitorBuilder {
	return &QueueMonitorBuilder{
		monitor: &QueueMonitor{
			conn:            conn,
			logger:          zap.NewNop(),
			mu:              sync.RWMutex{},
			pollingInterval: 1 * time.Second,
			monitored:       make([]*Queue, 0),
			metrics:         make(QueueMetrics),
			stopChan:        make(chan struct{}),
			handlers:        make([]QueueMetricsHandler, 0),
		},
	}
}

func (b *QueueMonitorBuilder) WithQueueMetricsHandler(handlers ...QueueMetricsHandler) *QueueMonitorBuilder {
	b.monitor.handlers = append(b.monitor.handlers, handlers...)
	return b
}

func (b *QueueMonitorBuilder) WithLogger(logger *zap.Logger) *QueueMonitorBuilder {
	b.monitor.logger = logger
	return b
}

func (b *QueueMonitorBuilder) WithQueue(queue *Queue) *QueueMonitorBuilder {
	b.monitor.monitored = append(b.monitor.monitored, queue)
	return b
}

func (b *QueueMonitorBuilder) WithPollingInterval(interval time.Duration) *QueueMonitorBuilder {
	b.monitor.pollingInterval = interval
	return b
}

func (b *QueueMonitorBuilder) Build() *QueueMonitor {
	return b.monitor
}
