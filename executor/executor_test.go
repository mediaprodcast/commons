package executor

import (
	"context"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type MockCommandRunner struct {
	mock.Mock
}

func (m *MockCommandRunner) CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	argsSlice := m.Called(ctx, name, args)
	return argsSlice.Get(0).(*exec.Cmd)
}

func newLogger() *zap.Logger {
	return zap.NewNop()
}

func TestShakaExecutorRun(t *testing.T) {
	// given
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "mock_binary", []string{}).
		Return(exec.Command("echo", "hello world"))

	executor := &Executor{
		logger:        newLogger(),
		commandRunner: mockRunner,
	}

	// when
	err := executor.RunWithContext(ctx)

	// then
	require.NoError(t, err)
	mockRunner.AssertCalled(t, "CommandContext", ctx, "mock_binary", []string{})
}

func TestShakaExecutorRunError(t *testing.T) {
	// given
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "shaka-packager", []string{}).
		Return(exec.Command("false"))

	executor := &Executor{
		logger:        newLogger(),
		commandRunner: mockRunner,
	}

	// when
	err := executor.RunWithContext(ctx)

	// then
	require.Error(t, err)
	require.Contains(t, err.Error(), "exit status 1")

	mockRunner.AssertCalled(t, "CommandContext", ctx, mock.Anything, mock.Anything)
}

func TestShakaExecutorRunAsync(t *testing.T) {
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "mock_binary", []string{}).
		Return(exec.Command("echo", "hello world"))

	executor := &Executor{
		logger:        newLogger(),
		commandRunner: mockRunner,
	}

	// when
	chn := executor.RunAsync(ctx)

	// then
	require.NoError(t, <-chn)
	mockRunner.AssertCalled(t, "CommandContext", ctx, "mock_binary", []string{})
}

func TestShakaExecutorErrorAsync(t *testing.T) {
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "shaka-packager", []string{}).
		Return(exec.Command("false"))

	executor := &Executor{
		logger:        newLogger(),
		commandRunner: mockRunner,
	}

	// when
	chn := executor.RunAsync(ctx)

	// then

	err := <-chn
	require.Error(t, err)
	require.Contains(t, err.Error(), "exit status 1")

	mockRunner.AssertCalled(t, "CommandContext", ctx, mock.Anything, mock.Anything)
}
