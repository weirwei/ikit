// package zlog 集成 zap ，自定义了一些语法糖，使
package zlog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	AppName  string `yaml:"appName"`  // 应用名称，default: ""
	Stdout   bool   `yaml:"stdout"`   // 是否打印到控制台，default: false
	Level    string `yaml:"level"`    // 日志级别，default: info
	Log2File bool   `yaml:"log2File"` // 是否打印到文件，default: true
	Path     string `yaml:"path"`     // 日志文件路径，default: ./log
}

var lcnf *logConfig

var (
	SugaredLogger *zap.SugaredLogger
	ZapLogger     *zap.Logger
)

type logConfig struct {
	ZapLevel zapcore.Level
	//hookField HookFieldFunc

	// 以下变量仅对开发环境生效
	Stdout   bool
	Log2File bool
	Path     string

	// 缓冲区
	BufferSwitch        bool
	BufferSize          int
	BufferFlushInterval time.Duration
}

func newLog() *logConfig {
	return &logConfig{
		ZapLevel: zapcore.InfoLevel,
		//hookField: defaultHook,

		Stdout:   false,
		Log2File: true,
		Path:     "log/",

		//// 缓冲区，如果不配置默认使用以下配置
		//BufferSwitch:        true,
		//BufferSize:          256 * 1024, // 256kb
		//BufferFlushInterval: 5 * time.Second,
	}
}

func InitLog(config Config) {
	lcnf = newLog()
	lcnf.ZapLevel = getLogLevel(config.Level)
	lcnf.Stdout = config.Stdout
	if len(lcnf.Path) > 0 {
		lcnf.Path = config.Path
	}
	if !strings.HasSuffix(lcnf.Path, "/") {
		lcnf.Path = lcnf.Path + "/"
	}
	if config.Log2File {
		lcnf.Log2File = true
	}
	// 获取目录路径
	dir := filepath.Dir(lcnf.Path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 目录不存在，创建目录 (包括所有父目录)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic("create log directory error: " + err.Error())
		}
	}

	if SugaredLogger == nil {
		if ZapLogger == nil {
			var zapCore []zapcore.Core
			if lcnf.Stdout {
				zapCore = append(zapCore,
					zapcore.NewCore(
						getEncoder(),
						zapcore.AddSync(os.Stdout),
						lcnf.ZapLevel))
			}
			if config.Log2File {
				zapCore = append(zapCore,
					zapcore.NewCore(
						getEncoder(),
						getLogWriter(config.AppName, logTypeNormal),
						lcnf.ZapLevel))

				zapCore = append(zapCore,
					zapcore.NewCore(
						getEncoder(),
						getLogWriter(config.AppName, logTypeAbnormal),
						zapcore.WarnLevel))
			}
			core := zapcore.NewTee(zapCore...)
			ZapLogger = zap.New(core,
				zap.Fields(),
				zap.WithCaller(true),
				zap.Development())
		}
		SugaredLogger = ZapLogger.Sugar()
	}
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func getLogWriter(name, loggerType string) (ws zapcore.WriteSyncer) {
	var w io.Writer
	// 打印到 name.log[.wf] 中
	var err error
	date := time.Now().Format("20060102")
	filename := filepath.Join(strings.TrimSuffix(lcnf.Path, "/"), appendLogFileTail(name, date, loggerType))
	w, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic("open log file error: " + err.Error())
	}

	// 开启缓冲区
	ws = &zapcore.BufferedWriteSyncer{
		WS:    zapcore.AddSync(w),
		Clock: nil,
	}
	return ws
}

const (
	logTypeNormal   = "normal"
	logTypeAbnormal = "abnormal"
)

// genFilename 拼装完整文件名
func appendLogFileTail(appName, date, loggerType string) string {
	var tailFixed string
	switch loggerType {
	case logTypeNormal:
		tailFixed = ".log"
	case logTypeAbnormal:
		tailFixed = ".log.wf"
	default:
		tailFixed = ".log"
	}
	return fmt.Sprintf("%s-%s%s", appName, date, tailFixed)
}

func CloseLogger() {
	if SugaredLogger != nil {
		_ = SugaredLogger.Sync()
	}

	if ZapLogger != nil {
		_ = ZapLogger.Sync()
	}
}
