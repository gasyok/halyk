package logger

import (
	"sync"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

var (
	l    *zap.Logger
	once sync.Once
)

func SetLevel(level zapcore.Level) {
	once.Do(func() {
		cfg := zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(level)
		l = zap.Must(cfg.Build())
	})
}

func SetTestLevel(t *testing.T, level zapcore.Level) {
	once.Do(func() {
		l = zaptest.NewLogger(t, zaptest.Level(level))
	})
}

func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	l.Fatal(msg, fields...)
}
