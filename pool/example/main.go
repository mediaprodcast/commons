package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/mediaprodcast/commons/pool"
	"go.uber.org/zap"
)

type CustomTask struct {
	Attempt int
}

func (t *CustomTask) Perform(ctx context.Context) error {
	t.Attempt++
	if t.Attempt < 3 {
		return errors.New("temporary gateway error")
	}

	// long running task
	time.Sleep(1 * time.Second)

	return nil
}

func (t *CustomTask) MaxAttempts() int {
	return 5
}

func (t *CustomTask) Name() string {
	return "task-id"
}

func (t *CustomTask) Backoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// For targeting 80% utilization
	workers := int(float64(runtime.NumCPU()) * 0.8)

	wp, _ := pool.New(
		pool.WithPoolConcurrency(workers),
		pool.WithPoolLogger(logger),
		pool.WithPoolName("custom"),
	)

	defer wp.Release()

	// Submit retryable task
	task := &CustomTask{}
	if err := wp.Submit(task); err != nil {
		logger.Error("failed to submit task", zap.Error(err))
	}

	if _, errs := wp.Wait(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err.Name(), err.Error())
		}
	}
}
