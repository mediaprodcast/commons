package executor

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"

	"github.com/shirou/gopsutil/v3/process"
)

// Constants for configuration and magic numbers.
const (
	bytesPerMB      = 1024 * 1024
	timeFormat      = "2006-01-02 15:04:05"
	metricsTemplate = `
----------------------------------------------------
PID: %d | CPU: %.2f%% | CPU PER CORE: %.2f%% | Mem: %d MB | Threads: %d | Status: %s
Start Time: %s | PPID: %d | User: %s
Exe Path: %s
----------------------------------------------------
`
)

// ExecuterMetrics struct holds process metrics and a logger.
type ExecuterMetrics struct {
	PID        int
	CPU        float64
	CPUPerCore float64
	Memory     uint64
	Threads    int32
	Status     string
	CreateTime int64
	PPID       int32
	Username   string
	ExePath    string
	logger     *zap.Logger // Updated to use zap.Logger
}

// NewExecuterMetrics initializes a new ExecuterMetrics instance with a zap logger.
func NewExecuterMetrics(pid int, logger *zap.Logger) *ExecuterMetrics {
	return &ExecuterMetrics{
		PID:    pid,
		logger: logger,
	}
}

// UpdateMetrics updates all available process metrics, logging individual errors.
func (pm *ExecuterMetrics) UpdateMetrics() error {
	proc, err := process.NewProcess(int32(pm.PID))
	if err != nil {
		pm.logger.Error(fmt.Sprintf("error getting process (PID: %d)", pm.PID), zap.Error(err))
		return fmt.Errorf("error getting process (PID: %d): %w", pm.PID, err)
	}

	pm.updateCPU(proc)
	pm.updateMemory(proc)
	pm.updateThreads(proc)
	pm.updateStatus(proc)
	pm.updateCreateTime(proc)
	pm.updatePPID(proc)
	pm.updateUsername(proc)
	pm.updateExePath(proc)

	return nil
}

// updateCPU updates the CPU usage metric.
func (pm *ExecuterMetrics) updateCPU(proc *process.Process) {
	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		pm.logger.Warn("failed to get CPU usage", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}

	numCPU := float64(runtime.NumCPU()) // Get total number of cores
	// To normalize CPU usage to a single core, divide by the number of cores:
	normalizedCPU := cpuPercent / numCPU

	pm.CPU = cpuPercent
	pm.CPUPerCore = normalizedCPU
}

// updateMemory updates the memory usage metric.
func (pm *ExecuterMetrics) updateMemory(proc *process.Process) {
	memInfo, err := proc.MemoryInfo()
	if err != nil {
		pm.logger.Warn("failed to get memory info", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.Memory = memInfo.RSS / bytesPerMB
}

// updateThreads updates the thread count metric.
func (pm *ExecuterMetrics) updateThreads(proc *process.Process) {
	threads, err := proc.NumThreads()
	if err != nil {
		pm.logger.Warn("failed to get thread count", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.Threads = threads
}

// updateStatus updates the process status metric.
func (pm *ExecuterMetrics) updateStatus(proc *process.Process) {
	status, err := proc.Status()
	if err != nil || len(status) == 0 {
		pm.logger.Warn("failed to get status", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.Status = status[0]
}

// updateCreateTime updates the process creation time metric.
func (pm *ExecuterMetrics) updateCreateTime(proc *process.Process) {
	createTime, err := proc.CreateTime()
	if err != nil {
		pm.logger.Warn("failed to get creation time", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.CreateTime = createTime
}

// updatePPID updates the parent process ID metric.
func (pm *ExecuterMetrics) updatePPID(proc *process.Process) {
	ppid, err := proc.Ppid()
	if err != nil {
		pm.logger.Warn("failed to get PPID", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.PPID = ppid
}

// updateUsername updates the username metric.
func (pm *ExecuterMetrics) updateUsername(proc *process.Process) {
	username, err := proc.Username()
	if err != nil {
		pm.logger.Warn("failed to get username", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.Username = username
}

// updateExePath updates the executable path metric.
func (pm *ExecuterMetrics) updateExePath(proc *process.Process) {
	exePath, err := proc.Exe()
	if err != nil {
		pm.logger.Warn("failed to get executable path", zap.Int("PID", pm.PID), zap.Error(err))
		return
	}
	pm.ExePath = exePath
}

// FormatMetrics returns a formatted string of the current metrics.
func (pm *ExecuterMetrics) FormatMetrics() string {
	createTime := time.Unix(0, pm.CreateTime*int64(time.Millisecond)).Format(timeFormat)
	return fmt.Sprintf(metricsTemplate,
		pm.PID,
		pm.CPU,
		pm.CPUPerCore,
		pm.Memory,
		pm.Threads,
		pm.Status,
		createTime,
		pm.PPID,
		pm.Username,
		pm.ExePath,
	)
}

// PrintMetrics continuously prints metrics at intervals until the context is cancelled.
func (pm *ExecuterMetrics) PrintMetrics(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			pm.logger.Info("stopping metrics collection", zap.Int("PID", pm.PID))
			return
		case <-ticker.C:
			if err := pm.UpdateMetrics(); err != nil {
				pm.logger.Error("error updating metrics", zap.Int("PID", pm.PID), zap.Error(err))
				return
			}
			fmt.Println(pm.FormatMetrics())
		}
	}
}
