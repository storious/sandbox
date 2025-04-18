package main

import (
	"fmt"

	"sandbox/config"
	"sandbox/internal/runtime"
	"sandbox/utils"
)

func main() {

	cfg := config.GetConfig()

	// initial logger
	utils.InitLog(cfg.LogPath, cfg.Debug)

	if ok := utils.Validate(cfg); !ok {
		fmt.Println("cpu", cfg.MaxCPUTime, "real", cfg.MaxRealTime, "mem", cfg.MaxMemory)
		panic("input args invalid")
	}
	// add runner
	runner := runtime.NewRunner(cfg)
	ret := runner.Run()

	// print result
	fmt.Printf(`{
		"cpu_time": %d,
		"real_time": %d,
		"memory": %d,
		"signal": %d,
		"exit_code": %d,
		"error": %d,
		"result": %d
	}`, ret.CPUTime, ret.RealTime, ret.Memory, ret.Signal, ret.ExitCode, ret.Error, ret.Result)
}
