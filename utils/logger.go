package utils

import (
	"go.uber.org/zap/zapcore"
	"strings"

	"go.uber.org/zap"
)

func NewLogger(logLevel string) *zap.Logger {
	var level zapcore.Level
	switch strings.ToUpper(logLevel) {
	case "WARN", "WARNING":
		level = zap.WarnLevel
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "ERROR":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	return l
}
