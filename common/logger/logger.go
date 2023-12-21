package logger

import (
	"strings"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var zapLevel *zap.AtomicLevel
var logInstance *zap.Logger

func NewLogger(logLevel string) *zap.Logger {
	var err error
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

	cfg := zap.Config{
		Level:    zap.NewAtomicLevelAt(level),
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "T",
			LevelKey:    "L",
			NameKey:     "N",
			CallerKey:   "C",
			MessageKey:  "M",
			LineEnding:  zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	if level == zap.DebugLevel {
		cfg.Development = false // add support for dev level
		cfg.EncoderConfig.StacktraceKey = "S"
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}

	logInstance, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	zapLevel = &cfg.Level

	return logInstance
}

func GetLogger() *zap.Logger {
	return logInstance
}

func SetLevel(logLevel string) {
	switch strings.ToUpper(logLevel) {
	case "WARN", "WARNING":
		zapLevel.SetLevel(zap.WarnLevel)
	case "DEBUG":
		zapLevel.SetLevel(zap.DebugLevel)
	case "INFO":
		zapLevel.SetLevel(zap.InfoLevel)
	case "ERROR":
		zapLevel.SetLevel(zap.ErrorLevel)
	}
}

func Sync() {
	logInstance.Sync() //nolint: errcheck
}
