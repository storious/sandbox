package utils

import (
	"log/slog"
	"os"
	"sandbox/config"
)

type logFunc func(msg string, args ...any)

var (
	Error logFunc
	Warn  logFunc
	Info  logFunc
	Debug logFunc
)

func Init() {
	var (
		target   = os.Stderr
		logLevel = &slog.LevelVar{}
		opt      = &slog.HandlerOptions{
			AddSource: true,
			Level:     logLevel,
		}
	)
	if config.DEBUG {
		logLevel.Set(slog.LevelDebug)
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
