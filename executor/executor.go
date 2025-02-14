package executor

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"syscall"

	"go.uber.org/zap"
)

var (
	ErrBinaryNotSet          = errors.New("binary path not configured")
	ErrCommandAlreadyRunning = errors.New("command is already running")
	ErrNoProcess             = errors.New("no process available")
)

type CommandRunner interface {
	CommandContext(context.Context, string, ...string) *exec.Cmd
}

type NativeCommandRunner struct{}

func (n *NativeCommandRunner) CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, args...)
}

type StdHandler = func(std io.Reader) error

type Executor struct {
	binaryPath    string
	logger        *zap.Logger
	args          []string
	commandRunner CommandRunner
	stdoutHandler StdHandler
	stderrHandler StdHandler

	cmdLock   sync.Mutex
	cmd       *exec.Cmd
	isRunning bool
}

type Option func(*Executor)

func NewExecutor(binaryPath string, logger *zap.Logger, opts ...Option) (*Executor, error) {
	if binaryPath == "" {
		return nil, ErrBinaryNotSet
	}

	e := &Executor{
		binaryPath:    binaryPath,
		logger:        logger,
		commandRunner: &NativeCommandRunner{},
	}

	for _, opt := range opts {
		opt(e)
	}

	return e, nil
}

func WithArgs(args ...string) Option {
	return func(e *Executor) {
		e.args = args
	}
}

func WithCommandRunner(runner CommandRunner) Option {
	return func(e *Executor) {
		e.commandRunner = runner
	}
}

func WithStdoutHandler(handler StdHandler) Option {
	return func(e *Executor) {
		e.stdoutHandler = handler
	}
}

func WithStderrHandler(handler StdHandler) Option {
	return func(e *Executor) {
		e.stderrHandler = handler
	}
}

func (e *Executor) PreviewCommand() (string, error) {
	if e.binaryPath == "" {
		return "", ErrBinaryNotSet
	}
	return fmt.Sprintf("%s %s", e.binaryPath, strings.Join(e.args, " ")), nil
}

func (e *Executor) RunWithContext(ctx context.Context) error {
	return e.execute(ctx)
}

func (e *Executor) Run() error {
	return e.RunWithContext(context.Background())
}

func (e *Executor) RunAsync(ctx context.Context) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		defer close(errCh)
		errCh <- e.execute(ctx)
	}()
	return errCh
}

func (e *Executor) Pause() error {
	e.cmdLock.Lock()
	defer e.cmdLock.Unlock()

	if e.cmd == nil || e.cmd.Process == nil {
		return ErrNoProcess
	}

	if err := e.cmd.Process.Signal(syscall.SIGSTOP); err != nil {
		e.logger.Error("Failed to pause process", zap.Error(err))
		return err
	}

	e.logger.Info("Process paused")
	return nil
}

func (e *Executor) Resume() error {
	e.cmdLock.Lock()
	defer e.cmdLock.Unlock()

	if e.cmd == nil || e.cmd.Process == nil {
		return ErrNoProcess
	}

	if err := e.cmd.Process.Signal(syscall.SIGCONT); err != nil {
		e.logger.Error("Failed to resume process", zap.Error(err))
		return err
	}

	e.logger.Info("Process resumed")
	return nil
}

func (e *Executor) Stop() error {
	e.cmdLock.Lock()
	defer e.cmdLock.Unlock()

	if e.cmd == nil || e.cmd.Process == nil {
		return ErrNoProcess
	}

	if err := e.cmd.Process.Kill(); err != nil {
		e.logger.Error("Failed to stop process", zap.Error(err))
		return err
	}

	e.logger.Info("Process stopped")
	return nil
}

func (e *Executor) execute(ctx context.Context) error {
	e.cmdLock.Lock()
	if e.isRunning {
		e.cmdLock.Unlock()
		return ErrCommandAlreadyRunning
	}
	e.isRunning = true
	defer func() {
		e.cmdLock.Lock()
		e.isRunning = false
		e.cmd = nil
		e.cmdLock.Unlock()
	}()

	cmd := e.commandRunner.CommandContext(ctx, e.binaryPath, e.args...)
	e.cmd = cmd
	e.cmdLock.Unlock()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("stdout pipe error: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("stderr pipe error: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("command start error: %w", err)
	}

	// Metrics
	// metrics := NewExecuterMetrics(cmd.Process.Pid, e.logger)
	// go metrics.PrintMetrics(context.Background(), 2*time.Second)

	var wg sync.WaitGroup
	handleStream := func(handler StdHandler, stream io.Reader, defaultHandler func(io.Reader)) {
		defer wg.Done()
		if handler != nil {
			if err := handler(stream); err != nil {
				e.logger.Error("stream handler error", zap.Error(err))
			}
		} else {
			defaultHandler(stream)
		}
	}

	wg.Add(2)
	go handleStream(e.stdoutHandler, stdout, func(r io.Reader) {
		readLog(r, func(line string) { fmt.Println(line) })
	})
	go handleStream(e.stderrHandler, stderr, func(r io.Reader) {
		readLog(r, func(line string) { fmt.Println(line) })
	})

	go func() {
		wg.Wait()
	}()

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("command wait error: %w", err)
	}

	return nil
}

func readLog(r io.Reader, callback func(string)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		callback(scanner.Text())
	}
}
