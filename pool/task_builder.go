package pool

import (
	"context"
	"errors"
	"math"
	"time"
)

type taskBuilder struct {
	handler     func(ctx context.Context) error
	maxAttempts int
	backoffTime time.Duration
	name        string
}

// Ensure taskBuilder implements the Task interface
var _ Task = (*taskBuilder)(nil)

type taskOption func(*taskBuilder)

func WithHandler(handler func(ctx context.Context) error) taskOption {
	return func(t *taskBuilder) {
		t.handler = handler
	}
}

func WithBackoff(backoffTime time.Duration) taskOption {
	return func(t *taskBuilder) {
		t.backoffTime = backoffTime
	}
}

func WithMaxAttempts(maxAttempts int) taskOption {
	return func(t *taskBuilder) {
		t.maxAttempts = maxAttempts
	}
}

func WithTaskName(name string) taskOption {
	return func(t *taskBuilder) {
		t.name = name
	}
}

func NewTask(opts ...taskOption) (Task, error) {
	t := &taskBuilder{
		maxAttempts: 1, // Default max attempts
	}
	for _, opt := range opts {
		opt(t)
	}

	if t.handler == nil {
		return nil, errors.New("task handler not defined")
	}

	return t, nil
}

func (t *taskBuilder) Perform(ctx context.Context) error {
	return t.handler(ctx)
}

func (t *taskBuilder) MaxAttempts() int {
	return t.maxAttempts
}

func (t *taskBuilder) Backoff(attempt int) time.Duration {
	if t.backoffTime > 0 {
		return t.backoffTime
	}

	// Start backoff at 2^0 * second, then 2^1, etc.
	return time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
}

func (t *taskBuilder) Name() string {
	return t.name
}
