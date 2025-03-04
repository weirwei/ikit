package zlog

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	ctxSugaredLogger = "ctxSugaredLogger"

	CtxKeyLogId     = "x-log-id"
	CtxKeyRequestId = "x-request-id"

	HeaderKeyLogId = "X-Log-Id"
)

func Info(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Info(args...)
		return
	}
	SugaredLogger.Info(args...)
}

func Infof(ctx *gin.Context, format string, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Infof(format, args...)
		return
	}
	SugaredLogger.Infof(format, args...)
}

func Debug(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Debug(args...)
		return
	}
	SugaredLogger.Debug(args...)
}

func Debugf(ctx *gin.Context, format string, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Debugf(format, args...)
		return
	}
	SugaredLogger.Debugf(format, args...)
}

func Error(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Error(args...)
		return
	}
	SugaredLogger.Error(args...)
}

func Errorf(ctx *gin.Context, format string, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Errorf(format, args...)
		return
	}
	SugaredLogger.Errorf(format, args...)
}

func Warn(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Warn(args...)
		return
	}
	SugaredLogger.Warn(args...)
}

func Warnf(ctx *gin.Context, format string, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Warnf(format, args...)
		return
	}
	SugaredLogger.Warnf(format, args...)
}

func Fatal(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Fatal(args...)
		return
	}
	SugaredLogger.Fatal(args...)
}

func Fatalf(ctx *gin.Context, format string, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Fatalf(format, args...)
		return
	}
	SugaredLogger.Fatalf(format, args...)
}

func Panic(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Panic(args...)
		return
	}
	SugaredLogger.Panic(args...)
}

func Panicf(ctx *gin.Context, format string, args ...interface{}) {
	if ctx != nil {
		newSugaredLogger(ctx).Panicf(format, args...)
		return
	}
	SugaredLogger.Panicf(format, args...)
}

func newSugaredLogger(ctx *gin.Context) *zap.SugaredLogger {
	if ctx == nil {
		return SugaredLogger
	}

	if t, exist := ctx.Get(ctxSugaredLogger); exist {
		if s, ok := t.(*zap.SugaredLogger); ok {
			return s
		}
	}

	s := SugaredLogger.With(
		zap.String("logId", GetLogID(ctx)),
		zap.String("uri", GetRequestUri(ctx)),
	)

	ctx.Set(ctxSugaredLogger, s)
	return s
}

func GetLogID(ctx *gin.Context) string {
	if ctx == nil {
		return genLogId()
	}
	if logId, exist := ctx.Get(CtxKeyLogId); exist {
		if s, ok := logId.(string); ok {
			return s
		}
	}
	var logID string
	if ctx.Request != nil && ctx.Request.Header != nil {
		logID = ctx.GetHeader(HeaderKeyLogId)
	}
	if logID == "" {
		logID = genLogId()
	}
	ctx.Set(CtxKeyLogId, logID)
	return logID
}

func GetRequestUri(ctx *gin.Context) string {
	if ctx == nil || ctx.Request == nil || ctx.Request.URL == nil {
		return ""
	}
	return ctx.Request.URL.Path
}

func genLogId() string {
	usec := uint64(time.Now().UnixNano())
	// 保证requestId不超过31位
	requestId := strconv.FormatUint(usec&0x7FFFFFFF|0x80000000, 10)
	return requestId
}
