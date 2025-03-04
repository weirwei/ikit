package zlog

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestSugar(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "www.baidu.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func TestInit(t *testing.T) {
	InitLog(Config{
		AppName:  "test",
		Level:    "debug",
		Stdout:   true,
		Log2File: true,
		Path:     "log/",
	})
	SugaredLogger.Info("test1")
	SugaredLogger.Warn("warn")
	SugaredLogger.Errorf("ERROR: %s", "test")

	CloseLogger()
}

func TestSugared(t *testing.T) {
	ctx := &gin.Context{}
	InitLog(Config{
		AppName:  "test",
		Level:    "debug",
		Stdout:   true,
		Log2File: true,
		Path:     "log/",
	})
	Info(ctx, "test1")
	Warn(ctx, "warn")
	Errorf(ctx, "ERROR: %s", "test")

	CloseLogger()
}
