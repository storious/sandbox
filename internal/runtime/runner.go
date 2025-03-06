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

// 假设的run函数，用于模拟执行配置
func (*Runner) Run() *config.Result {
	// 这里应该是执行外部程序的逻辑，以及收集结果的逻辑
	// 由于Go语言没有直接的方式来设置资源限制，这部分可能需要调用外部工具或cgo
	result := defaultResult()

	// check root
	uid := syscall.Getuid()
	if uid != 0 {
		// TODO
		panic("Operation not permitted: Runner need root privilege")
	}

	// check args
	return result
}
