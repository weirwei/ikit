package zlog

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogLevel(lv string) zapcore.Level {
	var level zapcore.Level
	switch strings.ToUpper(lv) {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "FATAL":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return level
}
