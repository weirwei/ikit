package zlog

import (
	"os"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestInitLog(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		check  func(t *testing.T)
	}{
		{
			name: "default config",
			config: Config{
				AppName:  "test",
				Stdout:   false,
				Level:    "info",
				Log2File: true,
				Path:     "./test_log",
			},
			check: func(t *testing.T) {
				if SugaredLogger == nil {
					t.Error("SugaredLogger should not be nil")
				}
				if ZapLogger == nil {
					t.Error("ZapLogger should not be nil")
				}
				if lcnf.ZapLevel != zapcore.InfoLevel {
					t.Errorf("expected log level to be %v, got %v", zapcore.InfoLevel, lcnf.ZapLevel)
				}
				if lcnf.Stdout {
					t.Error("expected Stdout to be false")
				}
				if !lcnf.Log2File {
					t.Error("expected Log2File to be true")
				}
				// 检查日志目录是否创建
				if _, err := os.Stat("./test_log"); err != nil {
					t.Errorf("log directory should exist: %v", err)
				}
			},
		},
		{
			name: "custom level config",
			config: Config{
				AppName:  "test",
				Stdout:   true,
				Level:    "debug",
				Log2File: true,
				Path:     "./test_log2",
			},
			check: func(t *testing.T) {
				if SugaredLogger == nil {
					t.Error("SugaredLogger should not be nil")
				}
				if ZapLogger == nil {
					t.Error("ZapLogger should not be nil")
				}
				if lcnf.ZapLevel != zapcore.DebugLevel {
					t.Errorf("expected log level to be %v, got %v", zapcore.DebugLevel, lcnf.ZapLevel)
				}
				if !lcnf.Stdout {
					t.Error("expected Stdout to be true")
				}
				if !lcnf.Log2File {
					t.Error("expected Log2File to be true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLog(tt.config)
			tt.check(t)
			CloseLogger()
			// 清理测试目录
			_ = os.RemoveAll(tt.config.Path)
		})
	}
}

func TestGetLogLevel(t *testing.T) {
	tests := []struct {
		level    string
		expected zapcore.Level
	}{
		{"debug", zapcore.DebugLevel},
		{"DEBUG", zapcore.DebugLevel},
		{"info", zapcore.InfoLevel},
		{"INFO", zapcore.InfoLevel},
		{"warn", zapcore.WarnLevel},
		{"WARN", zapcore.WarnLevel},
		{"error", zapcore.ErrorLevel},
		{"ERROR", zapcore.ErrorLevel},
		{"fatal", zapcore.FatalLevel},
		{"FATAL", zapcore.FatalLevel},
		{"invalid", zapcore.InfoLevel}, // 默认级别
		{"", zapcore.InfoLevel},        // 默认级别
	}

	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			level := getLogLevel(tt.level)
			if level != tt.expected {
				t.Errorf("getLogLevel(%q) = %v, want %v", tt.level, level, tt.expected)
			}
		})
	}
}

func TestAppendLogFileTail(t *testing.T) {
	tests := []struct {
		appName    string
		date       string
		loggerType string
		expected   string
	}{
		{
			appName:    "test",
			date:       "20240304",
			loggerType: logTypeNormal,
			expected:   "test-20240304.log",
		},
		{
			appName:    "test",
			date:       "20240304",
			loggerType: logTypeAbnormal,
			expected:   "test-20240304.log.wf",
		},
		{
			appName:    "test",
			date:       "20240304",
			loggerType: "unknown",
			expected:   "test-20240304.log",
		},
	}

	for _, tt := range tests {
		t.Run(tt.appName+"-"+tt.loggerType, func(t *testing.T) {
			result := appendLogFileTail(tt.appName, tt.date, tt.loggerType)
			if result != tt.expected {
				t.Errorf("appendLogFileTail(%q, %q, %q) = %v, want %v",
					tt.appName, tt.date, tt.loggerType, result, tt.expected)
			}
		})
	}
}

func TestCloseLogger(t *testing.T) {
	// 初始化日志
	config := Config{
		AppName:  "test",
		Stdout:   true,
		Level:    "info",
		Log2File: true,
		Path:     "./test_log",
	}
	InitLog(config)

	// 确保日志器已初始化
	if SugaredLogger == nil {
		t.Error("SugaredLogger should not be nil")
	}
	if ZapLogger == nil {
		t.Error("ZapLogger should not be nil")
	}

	// 关闭日志器
	CloseLogger()

	// 清理测试目录
	_ = os.RemoveAll(config.Path)
}
