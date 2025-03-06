package config

const UNLIMITED = -1

// 定义配置和结果结构体
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
	Env                  []string
	LogPath              string
	SeccompRuleName      string
	UID                  uint32
	GID                  uint32
}

const (
	SUCCESS = -iota
	INVALID_CONFIG
	FORK_FAILED
	PTHREAD_FAILED
	WAIT_FAILED
	ROOT_REQUIRED
	LOAD_SECCOMP_FAILED
	SETRLIMIT_FAILED
	DUP2_FAILED
	SETUID_FAILED
	EXECVE_FAILED
	SPJ_ERROR
)

type Result struct {
	CPUTime  int
	RealTime int
	Memory   int64
	Signal   int
	ExitCode int
	Error    int
	Result   int
}

const (
	WRONG_ANSWER = iota + 200
	CPU_TIME_LIMIT_EXCEEDED
	REAL_TIME_LIMIT_EXCEEDED
	MEMORY_LIMIT_EXCEEDED
	RUNTIME_ERROR
	SYSTEM_ERROR
)
