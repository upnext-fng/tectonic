package logger

import (
	"upnext-fng/tgpl-api/static"

	"go.uber.org/zap/zapcore"
)

type Option func(*Logger)

func WithDevelopment(isDevelop bool) Option {
	return func(l *Logger) {
		l.isDevelopment = isDevelop
	}
}

func WithLevel(level string) Option {
	logLevel := zapcore.InfoLevel

	if lvl, exist := static.LogLevels[level]; exist {
		logLevel = lvl
	}

	return func(l *Logger) {
		l.level = logLevel
	}
}
