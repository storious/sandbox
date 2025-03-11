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

func (*Runner) Run() *config.Result {
	result := defaultResult()

	// check root
	uid := syscall.Getuid()
	if uid != 0 {
		// TODO use virtual cotianer replace
		panic("Operation forbid: Runner need root privilege")
	}

	// check args
	return result
}
