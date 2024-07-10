package ilog

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var (
	debugLog = log.New(os.Stdout, fmt.Sprintf("%s[debug]%s", lightGreen, none), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	infoLog  = log.New(os.Stdout, fmt.Sprintf("%s[info ]%s", lightPurple, none), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	warnLog  = log.New(os.Stdout, fmt.Sprintf("%s[warn ]%s", yellow, none), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	errorLog = log.New(os.Stdout, fmt.Sprintf("%s[error]%s", lightRed, none), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	fatalLog = log.New(os.Stdout, fmt.Sprintf("%s[fatal]%s", lightRed, none), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	panicLog = log.New(os.Stdout, fmt.Sprintf("%s[painc]%s", lightRed, none), log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	loggers  = []*log.Logger{debugLog, infoLog, warnLog, errorLog, fatalLog, panicLog}
	mu       sync.Mutex
)

// SetLevel 设置日志等级，打印指定等级以上的日志。默认打印全部
func SetLevel(level Level) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if LevelDebug > level {
		debugLog.SetOutput(io.Discard)
	}
	if LevelInfo > level {
		infoLog.SetOutput(io.Discard)
	}
	if LevelWarn > level {
		warnLog.SetOutput(io.Discard)
	}
	if LevelError > level {
		errorLog.SetOutput(io.Discard)
	}
	if LevelFatal > level {
		fatalLog.SetOutput(io.Discard)
	}
	if LevelPanic > level {
		panicLog.SetOutput(io.Discard)
	}
}
