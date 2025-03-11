package config

// runner status code
const (
	SUCCESS             = 0
	INVALID_CONFIG      = 1
	FORK_FAILED         = 2
	PTHREAD_FAILED      = 3
	WAIT_FAILED         = 4
	ROOT_REQUIRED       = 5
	LOAD_SECCOMP_FAILED = 6
	SETRLIMIT_FAILED    = 7
	DUP2_FAILED         = 8
	SETUID_FAILED       = 9
	EXECVE_FAILED       = 10
	SPJ_ERROR           = 11
)

// judge status code
const (
	ACCEPT                   = 200
	WRONG_ANSWER             = 600
	CPU_TIME_LIMIT_EXCEEDED  = 601
	REAL_TIME_LIMIT_EXCEEDED = 602
	MEMORY_LIMIT_EXCEEDED    = 603
	RUNTIME_ERROR            = 604
	SYSTEM_ERROR             = 605
)

func StatusMessage(code int) (res string) {
	switch code {
	case 0:
		res = "Success"
	case 1:
		res = "Invalid Config"
	case 2:
		res = "Fork Failed"
	case 3:
		res = "Pthread Failed"
	case 4:
		res = "Wait Failed"
	case 5:
		res = "Root Required"
	case 6:
		res = "Load Seccomp Failed"
	case 7:
		res = "SetrLimit Failed"
	case 8:
		res = "Duplicate Failed"
	case 9:
		res = "Set Uid Failed"
	case 10:
		res = "Execve Failed"
	case 11:
		res = "Special Judge Error"
	case 200:
		res = "Accept"
	case 600:
		res = "Wrong Answer"
	case 601:
	case 602:
		res = "Time Limit Exceeded"
	case 603:
		res = "Memory Limit Exceeded"
	case 604:
		res = "Runtime Error"
	case 605:
		res = "System Error"
	default:
		res = "Not Found the Status code"
	}
	return
}
