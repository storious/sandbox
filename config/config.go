package config

import (
	"flag"
	"fmt"
	"os"
)

// define config and result struct

const UNLIMITED = -1

type Config struct {
	MaxCPUTime           int   `validate:"required"`
	MaxRealTime          int   `validate:"required"`
	MaxMemory            int64 `validate:"required"`
	MemoryLimitCheckOnly bool
	MaxStack             int64  `validate:"required,gte=1"`
	MaxProcessNumber     int    `validate:"required,gte=1"`
	MaxOutputSize        int64  `validate:"required,gte=1"`
	ExePath              string `validate:"required"`
	InputPath            string `validate:"required"`
	OutputPath           string `validate:"required"`
	ErrorPath            string `validate:"required"`
	Args                 []string
	Env                  string
	LogPath              string
	SeccompRuleName      string
	UID                  int
	GID                  int
	Debug                bool
}

type Result struct {
	CPUTime  int
	RealTime int
	Memory   int64
	Signal   int
	ExitCode int
	Error    int
	Result   int
}

func GetConfig() *Config {

	// define command line args
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
	uid := flag.Int("uid", 65534, "UID (default 65534)")
	gid := flag.Int("gid", 65534, "GID (default 65534)")
	Debug := flag.Bool("Debug", false, "DEBUG mode (default false)")

	flag.Parse()

	if *Debug {
		fmt.Println("Running in Deubg mode")
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		fmt.Println("Version: 1.0.0")
		os.Exit(0)
	}

	// build config
	return &Config{
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
		Env:                  *env,
		LogPath:              *logPath,
		SeccompRuleName:      *seccompRuleName,
		UID:                  *uid,
		GID:                  *gid,
		Debug:                *Debug,
	}
}
