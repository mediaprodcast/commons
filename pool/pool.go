package pool

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type RetryTask interface {
	Task
	MaxAttempts() int
	Backoff(attempt int) time.Duration
}

type Task interface {
	Name() string
	Perform(ctx context.Context) error
}

type Error struct {
	err  string
	name string
}

func (e *Error) Error() string {
	return e.err
}

func (e *Error) Name() string {
	return e.name
}

type Metrics struct {
	StartTime   time.Time
	EndTime     time.Time
	CpuTimeUsed time.Duration
	WorkersUsed int
	TaskCount   int
}

type Pool struct {
	pool        *ants.Pool
	logger      *zap.Logger
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	poolLook    sync.Mutex
	concurrency int
	name        string
	errors      []*Error
	errLock     sync.Mutex
	// Metrics
	startTime      time.Time
	endTime        time.Time
	cpuTimeUsed    time.Duration
	maxWorkersUsed int
	taskCount      int
	metricsLock    sync.Mutex
}

func New(opts ...poolOption) (*Pool, error) {
	ctx, cancel := context.WithCancel(context.Background())

	p := &Pool{
		concurrency: 1,
		ctx:         ctx,
		cancel:      cancel,
		logger:      zap.NewNop(),
	}
	// Set Options
	for _, fn := range opts {
		fn(p)
	}

	// Logger with component
	p.logger = p.logger.With(
		zap.String("component", "worker_pool"),
		zap.String("name", p.name),
	)

	pool, err := ants.NewPool(
		p.concurrency,
		ants.WithPanicHandler(p.panicHandler),
		ants.WithLogger(antsLogger{logger: p.logger}),
	)

	if err != nil {
		cancel()
		return nil, err
	}

	p.pool = pool
	return p, nil
}

func (p *Pool) Submit(task Task) error {
	p.poolLook.Lock()
	defer p.poolLook.Unlock()

	p.wg.Add(1)
	return p.pool.Submit(func() {
		defer p.wg.Done()

		p.taskCount++

		taskStart := time.Now()

		// Update startTime
		p.metricsLock.Lock()
		if p.startTime.IsZero() || taskStart.Before(p.startTime) {
			p.startTime = taskStart
		}
		p.metricsLock.Unlock()

		// Update maxWorkersUsed
		currentWorkers := p.pool.Running()
		p.metricsLock.Lock()
		if currentWorkers > p.maxWorkersUsed {
			p.maxWorkersUsed = currentWorkers
		}
		p.metricsLock.Unlock()

		p.executeWithRetry(task)

		taskEnd := time.Now()
		p.metricsLock.Lock()
		if taskEnd.After(p.endTime) {
			p.endTime = taskEnd
		}
		p.metricsLock.Unlock()
	})
}

func (p *Pool) executeWithRetry(task Task) {
	var attempt int
	maxAttempts := 1
	backoffFunc := defaultBackoff

	if rt, ok := task.(RetryTask); ok {
		maxAttempts = rt.MaxAttempts()
		backoffFunc = rt.Backoff
	}

	var err error
	var totalTaskCpu time.Duration

	for attempt = 1; attempt <= maxAttempts; attempt++ {
		p.logger.Debug("started", zap.String("name", task.Name()))

		start := time.Now()
		err = task.Perform(p.ctx)
		duration := time.Since(start)
		totalTaskCpu += duration

		if err == nil {
			p.logTaskSuccess(task, attempt, duration)
			break
		}

		p.logTaskRetry(task, attempt, maxAttempts, err, duration)

		if attempt < maxAttempts {
			select {
			case <-time.After(backoffFunc(attempt)):
			case <-p.ctx.Done():
				p.logger.Warn("context canceled during backoff",
					zap.Int("attempt", attempt),
					zap.String("name", task.Name()),
				)
				return
			}
		}
	}

	p.metricsLock.Lock()
	p.cpuTimeUsed += totalTaskCpu
	p.metricsLock.Unlock()

	if err != nil {
		p.errLock.Lock()
		p.errors = append(p.errors, &Error{err: err.Error(), name: task.Name()})
		p.errLock.Unlock()
	}
}

// Release gracefully stops the pool
func (p *Pool) Release() {
	p.poolLook.Lock()
	defer p.poolLook.Unlock()

	p.cancel()       // Cancel context for running tasks
	p.pool.Release() // Release ants pool resources
}

// Wait for all tasks to complete
func (p *Pool) Wait() (*Metrics, []*Error) {
	p.wg.Wait()

	p.metricsLock.Lock()
	metrics := &Metrics{
		StartTime:   p.startTime,
		EndTime:     p.endTime,
		CpuTimeUsed: p.cpuTimeUsed,
		WorkersUsed: p.maxWorkersUsed,
		TaskCount:   p.taskCount,
	}
	p.metricsLock.Unlock()

	p.errLock.Lock()
	errors := make([]*Error, len(p.errors))
	copy(errors, p.errors)
	p.errLock.Unlock()

	p.logger.Debug("complete",
		zap.Int("errors", len(errors)),
		zap.Int("total_tasks", metrics.TaskCount),
		zap.Int("workers_used", metrics.WorkersUsed),
		zap.Duration("cpu_time_used", metrics.CpuTimeUsed),
		zap.String("start_time", metrics.StartTime.Format(time.RFC3339)),
		zap.String("end_time", metrics.EndTime.Format(time.RFC3339)),
		zap.String("elapsed_time", metrics.EndTime.Sub(metrics.StartTime).String()),
	)

	return metrics, errors
}

func defaultBackoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

// Logging helpers
func (p *Pool) logTaskSuccess(task Task, attempt int, duration time.Duration) {
	p.logger.Debug("succeeded",
		zap.Int("attempt", attempt),
		zap.Duration("duration", duration),
		zap.String("name", task.Name()),
	)
}

func (p *Pool) logTaskRetry(task Task, attempt, max int, err error, duration time.Duration) {
	fields := []zap.Field{
		zap.Int("attempt", attempt),
		zap.Int("max_attempts", max),
		zap.Duration("duration", duration),
		zap.String("name", task.Name()),
		zap.Error(err),
	}

	if attempt == max {
		p.logger.Error("task final attempt failed", fields...)
	} else {
		p.logger.Warn("task failed, will retry", fields...)
	}
}

func (p *Pool) panicHandler(r interface{}) {
	err := fmt.Errorf("panic occurred: %v", r)
	p.logger.Error("worker panic recovered",
		zap.Any("recovery_info", r),
		zap.Stack("stack"),
	)
	p.errLock.Lock()
	p.errors = append(p.errors, &Error{err: err.Error()})
	p.errLock.Unlock()
}

// Custom ants logger adapter
type antsLogger struct {
	logger *zap.Logger
}

func (l antsLogger) Printf(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}
