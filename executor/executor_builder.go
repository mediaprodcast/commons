package executor

import (
	"go.uber.org/zap"
)

type ExecutorBuilder struct {
	executor *Executor
}

func NewExecutorBuilder() *ExecutorBuilder {
	return &ExecutorBuilder{
		executor: &Executor{
			logger:        zap.NewNop(),
			args:          make([]string, 0),
			commandRunner: &NativeCommandRunner{},
		},
	}
}

func (e *ExecutorBuilder) WithLogger(logger *zap.Logger) *ExecutorBuilder {
	e.executor.logger = logger
	return e
}

func (e *ExecutorBuilder) WithBinaryPath(binaryPath string) *ExecutorBuilder {
	e.executor.binaryPath = binaryPath
	return e
}

func (e *ExecutorBuilder) WithArgs(args []string) *ExecutorBuilder {
	e.executor.args = append(e.executor.args, args...)
	return e
}

func (e *ExecutorBuilder) WithStdoutHandler(handler StdHandler) *ExecutorBuilder {
	e.executor.stdoutHandler = handler
	return e
}

func (e *ExecutorBuilder) WithStderrHandler(handler StdHandler) *ExecutorBuilder {
	e.executor.stderrHandler = handler
	return e
}

func (e *ExecutorBuilder) Build() *Executor {
	return e.executor
}
