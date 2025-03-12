package runtime

import (
	"context"
	"os"
	"os/exec"
	"sandbox/config"
	"sandbox/utils/logger"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func ChildProcess(cfg *config.Config) (duration int, status syscall.WaitStatus, rusage syscall.Rusage) {
	var (
		err    error
		input  *os.File
		output *os.File
		errput *os.File
		start  time.Time
		end    time.Time
	)

	// Set Stack limit
	if cfg.MaxStack != config.UNLIMITED {
		err = syscall.Setrlimit(syscall.RLIMIT_STACK, &syscall.Rlimit{
			Cur: uint64(cfg.MaxStack),
			Max: uint64(cfg.MaxStack),
		})
		if err != nil {
			logger.Error(config.StatusMessage(config.SETRLIMIT_FAILED) + ": Stack")
			return
		}
	}

	// Set Memory limit
	if !cfg.MemoryLimitCheckOnly {
		if cfg.MaxMemory != config.UNLIMITED {
			err = syscall.Setrlimit(syscall.RLIMIT_AS, &syscall.Rlimit{
				Cur: uint64(cfg.MaxMemory * 2),
				Max: uint64(cfg.MaxMemory * 2),
			})
			if err != nil {
				logger.Error(config.StatusMessage(config.SETRLIMIT_FAILED) + ": Memory")
				return
			}
		}
	}

	// Set CPU limit (seconds)
	if cfg.MaxCPUTime != config.UNLIMITED {
		err = syscall.Setrlimit(syscall.RLIMIT_CPU, &syscall.Rlimit{
			Cur: uint64(cfg.MaxMemory+1000) / 1000,
			Max: uint64(cfg.MaxMemory+1000) / 1000,
		})
		if err != nil {
			logger.Error(config.StatusMessage(config.SETRLIMIT_FAILED) + ": CPU Time")
			return
		}
	}

	// TODO repalce a standard method in future
	// Set Process number
	if cfg.MaxProcessNumber != config.UNLIMITED {
		err = unix.Setrlimit(unix.RLIMIT_NPROC, &unix.Rlimit{
			Cur: uint64(cfg.MaxMemory),
			Max: uint64(cfg.MaxMemory),
		})
		if err != nil {
			logger.Error(config.StatusMessage(config.SETRLIMIT_FAILED) + ": Process number")
			return
		}
	}

	// Set Output size
	if cfg.MaxOutputSize != config.UNLIMITED {
		err = syscall.Setrlimit(syscall.RLIMIT_FSIZE, &syscall.Rlimit{
			Cur: uint64(cfg.MaxMemory),
			Max: uint64(cfg.MaxMemory),
		})
		if err != nil {
			logger.Error(config.StatusMessage(config.SETRLIMIT_FAILED) + ": Output Size")
			return
		}
	}

	// Set input path
	if len(cfg.InputPath) > 0 {
		input, err = os.Open(cfg.InputPath)
		if err != nil {
			logger.Error(config.StatusMessage(config.DUP2_FAILED) + ": input path error")
			return
		}
		// redirect to stdin
		if err = syscall.Dup2(int(input.Fd()), int(os.Stdin.Fd())); err != nil {
			logger.Error(config.StatusMessage(config.DUP2_FAILED) + ": input redirect error")
			return
		}
	}
	// Set output path
	if len(cfg.OutputPath) > 0 {
		output, err = os.OpenFile(cfg.OutputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		if err != nil {
			logger.Error(config.StatusMessage(config.DUP2_FAILED) + ": output path error")
			return
		}
		// redirect to out file
		if err = syscall.Dup2(int(output.Fd()), int(os.Stdout.Fd())); err != nil {
			logger.Error(config.StatusMessage(config.DUP2_FAILED) + ": output redirect error")
			return
		}
	}

	// Set error path
	if len(cfg.ErrorPath) > 0 {
		if len(cfg.OutputPath) > 0 && cfg.OutputPath == cfg.ErrorPath {
			errput = output
		} else {
			errput, err = os.OpenFile(cfg.ErrorPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
			if err != nil {
				logger.Error(config.StatusMessage(config.DUP2_FAILED) + ": error path error")
				return
			}
		}
		// redirect to stderr
		if err = syscall.Dup2(int(errput.Fd()), int(os.Stderr.Fd())); err != nil {
			logger.Error(config.StatusMessage(config.DUP2_FAILED) + ": error redirect error")
			return
		}
	}

	// group id list
	groupIDs := []int{cfg.GID}
	if cfg.GID != -1 && (syscall.Setgid(cfg.GID) != nil || syscall.Setgroups(groupIDs) != nil) {
		logger.Error(config.StatusMessage(config.SETUID_FAILED) + ": set groups failed")
		return
	}

	// set uid
	if cfg.UID != -1 && syscall.Setuid(cfg.UID) != nil {
		logger.Error(config.StatusMessage(config.SETUID_FAILED) + ": set uid failed")
		return
	}

	// load seccomp
	// TODO ...

	// exec command with time out limit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MaxRealTime))
	defer cancel()

	start = time.Now()

	cmd := exec.CommandContext(ctx, cfg.ExePath, cfg.Args...)
	cmd.Env = strings.Split(cfg.Env, " ")
	if err = cmd.Start(); err != nil {
		logger.Error(config.StatusMessage(config.EXECVE_FAILED) + ": exec failed")
		return
	}

	if _, err = syscall.Wait4(cmd.Process.Pid, &status, syscall.WSTOPPED, &rusage); err != nil {
		cmd.Process.Kill()
		logger.Error("Wait4 command exec failed -> " + err.Error())
		if ctx.Err() == context.DeadlineExceeded {
			logger.Warn("command time out")
		}
	}
	end = time.Now()
	duration = int(end.Sub(start).Milliseconds())
	return
}
