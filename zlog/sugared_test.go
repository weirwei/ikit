package zlog

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/test", nil)
	return c, w
}

func TestLoggerWithContext(t *testing.T) {
	// 初始化日志配置
	config := Config{
		AppName:  "test",
		Stdout:   true,
		Level:    "debug",
		Log2File: false,
	}
	InitLog(config)

	ctx, _ := setupTestContext()

	// 测试各种日志级别
	t.Run("test all log levels with context", func(t *testing.T) {
		// Debug
		Debug(ctx, "debug message")
		Debugf(ctx, "debug message %s", "formatted")

		// Info
		Info(ctx, "info message")
		Infof(ctx, "info message %s", "formatted")

		// Warn
		Warn(ctx, "warn message")
		Warnf(ctx, "warn message %s", "formatted")

		// Error
		Error(ctx, "error message")
		Errorf(ctx, "error message %s", "formatted")

		// 不测试 Fatal 和 Panic，因为它们会终止程序
	})

	// 测试无 context 的情况
	t.Run("test all log levels without context", func(t *testing.T) {
		Debug(nil, "debug message")
		Debugf(nil, "debug message %s", "formatted")
		Info(nil, "info message")
		Infof(nil, "info message %s", "formatted")
		Warn(nil, "warn message")
		Warnf(nil, "warn message %s", "formatted")
		Error(nil, "error message")
		Errorf(nil, "error message %s", "formatted")
	})

	CloseLogger()
}

func TestGetLogID(t *testing.T) {
	ctx, _ := setupTestContext()

	t.Run("get log id from header", func(t *testing.T) {
		expectedLogID := "123456"
		ctx.Request.Header.Set(HeaderKeyLogId, expectedLogID)
		logID := GetLogID(ctx)
		if logID != expectedLogID {
			t.Errorf("GetLogID() = %v, want %v", logID, expectedLogID)
		}
	})

	t.Run("get log id from context", func(t *testing.T) {
		expectedLogID := "789012"
		ctx.Set(CtxKeyLogId, expectedLogID)
		logID := GetLogID(ctx)
		if logID != expectedLogID {
			t.Errorf("GetLogID() = %v, want %v", logID, expectedLogID)
		}
	})

	t.Run("generate new log id", func(t *testing.T) {
		ctx, _ := setupTestContext() // 使用新的 context
		logID := GetLogID(ctx)
		if logID == "" {
			t.Error("GetLogID() returned empty string")
		}
		if len(logID) != 10 {
			t.Errorf("GetLogID() returned id with length %d, want 10", len(logID))
		}
	})

	t.Run("get log id with nil context", func(t *testing.T) {
		logID := GetLogID(nil)
		if logID == "" {
			t.Error("GetLogID() returned empty string")
		}
		if len(logID) != 10 {
			t.Errorf("GetLogID() returned id with length %d, want 10", len(logID))
		}
	})
}

func TestGetRequestUri(t *testing.T) {
	t.Run("get uri from valid request", func(t *testing.T) {
		ctx, _ := setupTestContext()
		ctx.Request.URL.Path = "/test/path"
		uri := GetRequestUri(ctx)
		if uri != "/test/path" {
			t.Errorf("GetRequestUri() = %v, want /test/path", uri)
		}
	})

	t.Run("get uri with nil context", func(t *testing.T) {
		uri := GetRequestUri(nil)
		if uri != "" {
			t.Errorf("GetRequestUri() = %v, want empty string", uri)
		}
	})
}

func TestNewSugaredLogger(t *testing.T) {
	config := Config{
		AppName:  "test",
		Stdout:   true,
		Level:    "debug",
		Log2File: false,
	}
	InitLog(config)

	t.Run("new logger with context", func(t *testing.T) {
		ctx, _ := setupTestContext()
		logger := newSugaredLogger(ctx)
		if logger == nil {
			t.Error("newSugaredLogger() returned nil")
		}

		// 再次获取应该返回相同的 logger
		logger2 := newSugaredLogger(ctx)
		if logger != logger2 {
			t.Error("second call to newSugaredLogger() returned different logger")
		}
	})

	t.Run("new logger with nil context", func(t *testing.T) {
		logger := newSugaredLogger(nil)
		if logger == nil {
			t.Error("newSugaredLogger() returned nil")
		}
		if logger != SugaredLogger {
			t.Error("newSugaredLogger() with nil context should return global SugaredLogger")
		}
	})

	CloseLogger()
}
