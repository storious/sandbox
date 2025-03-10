package utils

import (
	"fmt"
	"log/slog"
	"os"
)

type logFunc func(msg string, args ...any)

var (
	Error logFunc
	Warn  logFunc
	Info  logFunc
	Debug logFunc
)

func InitLog(path string, debug bool) {
	var (
		target   = os.Stderr
		logLevel = &slog.LevelVar{}
		opt      = &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		}
		err error
	)
	if debug {
		logLevel.Set(slog.LevelDebug)
	}

	if len(path) != 0 && !debug {
		target, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0o666)
		if err != nil {
			fmt.Println("open log file error")
			return
		}
	}

	handler := slog.NewJSONHandler(target, opt)
	logger := slog.New(handler)

	// set logger
	{
		Error = logger.Error
		Warn = logger.Warn
		Info = logger.Info
		Debug = logger.Debug
	}
}
