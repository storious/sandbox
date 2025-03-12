package runtime

import (
	"sandbox/config"
	"syscall"
)

type Runner struct {
	*config.Config
}

func NewRunner(cfg *config.Config) *Runner {
	return &Runner{
		Config: cfg,
	}
}

func defaultResult() *config.Result {
	return &config.Result{
		Result:   config.SUCCESS,
		Error:    config.SUCCESS,
		CPUTime:  0,
		RealTime: 0,
		Signal:   0,
		ExitCode: 0,
		Memory:   0,
	}
}

func (r *Runner) Run() *config.Result {
	result := defaultResult()

	// check root
	uid := syscall.Getuid()
	if uid != 0 {
		panic("Operation forbid: Runner need root privilege")
	}

	duration, status, rusage := ChildProcess(r.Config)

	result.RealTime = duration
	if status.Signal() != 0 {
		result.Signal = int(status.Signal())
	}

	// TODO complete
	if status.Signal() == syscall.SIGUSR1 {
		result.Result = config.SYSTEM_ERROR
	} else {
		result.ExitCode = status.ExitStatus()
		result.CPUTime = int(rusage.Utime.Sec*1000 + rusage.Utime.Usec/1000)
		result.Memory = rusage.Maxrss

		if result.ExitCode != 0 {
			result.Result = config.RUNTIME_ERROR
		}

		if result.Signal == int(syscall.SIGSEGV) {
			if r.MaxMemory != config.UNLIMITED && result.Memory > r.MaxMemory {
				result.Result = config.MEMORY_LIMIT_EXCEEDED
			} else {
				result.Result = config.RUNTIME_ERROR
			}
		} else {
			if result.Signal != 0 {
				result.Result = config.RUNTIME_ERROR
			}
			if r.MaxMemory != config.UNLIMITED && result.Memory > r.MaxMemory {
				result.Result = config.MEMORY_LIMIT_EXCEEDED
			}
			if r.MaxRealTime != config.UNLIMITED && result.RealTime > r.MaxRealTime {
				result.Result = config.REAL_TIME_LIMIT_EXCEEDED
			}
			if r.MaxCPUTime != config.UNLIMITED && result.CPUTime > r.MaxCPUTime {
				result.Result = config.CPU_TIME_LIMIT_EXCEEDED
			}
		}
	}
	return result
}
