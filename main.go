package main

import (
	"flag"
	"fmt"

	"sandbox/config"
	"sandbox/internal/runtime"
	"sandbox/utils"
)

func main() {
	// 定义命令行参数
	help := flag.Bool("help", false, "Display This Help And Exit")
	version := flag.Bool("version", false, "Display Version Info And Exit")
	maxCPUTime := flag.Int("max_cpu_time", 1000, "Max CPU Time (ms)")
	maxRealTime := flag.Int("max_real_time", 1000, "Max Real Time (ms)")
	maxMemory := flag.Int64("max_memory", 16*1024*1024, "Max Memory (byte)")
	memoryLimitCheckOnly := flag.Bool("memory_limit_check_only", false, "only check memory usage, do not setrlimit (default False)")
	maxStack := flag.Int64("max_stack", 16*1024*1024, "Max Stack (byte, default 16M)")
	maxProcessNumber := flag.Int("max_process_number", 1, "Max Process Number")
	maxOutputSize := flag.Int64("max_output_size", 1, "Max Output Size (byte)")
	exePath := flag.String("exe_path", "./", "Exe Path")
	inputPath := flag.String("input_path", "/dev/stdin", "Input Path")
	outputPath := flag.String("output_path", "/dev/stdout", "Output Path")
	errorPath := flag.String("error_path", "/dev/stderr", "Error Path")
	args := flag.Args()
	env := flag.String("env", "", "Env")
	logPath := flag.String("log_path", "judger.log", "Log Path")
	seccompRuleName := flag.String("seccomp_rule_name", "", "Seccomp Rule Name")
	uid := flag.Uint("uid", 65534, "UID (default 65534)")
	gid := flag.Uint("gid", 65534, "GID (default 65534)")
	config.DEBUG = *flag.Bool("Debug", false, "DEBUG mode (default false)")

	flag.Parse()

	// initial logger
	utils.Init()

	if *help {
		flag.Usage()
		return
	}

	if *version {
		fmt.Println("Version: 1.0.0") // 示例版本号
		return
	}

	// build config
	cfg := &config.Config{
		MaxCPUTime:           *maxCPUTime,
		MaxRealTime:          *maxRealTime,
		MaxMemory:            *maxMemory,
		MemoryLimitCheckOnly: *memoryLimitCheckOnly,
		MaxStack:             *maxStack,
		MaxProcessNumber:     *maxProcessNumber,
		MaxOutputSize:        *maxOutputSize,
		ExePath:              *exePath,
		InputPath:            *inputPath,
		OutputPath:           *outputPath,
		ErrorPath:            *errorPath,
		Args:                 args,
		Env:                  []string{*env},
		LogPath:              *logPath,
		SeccompRuleName:      *seccompRuleName,
		UID:                  uint32(*uid),
		GID:                  uint32(*gid),
	}

	if ok := utils.Validate(cfg); !ok {
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
