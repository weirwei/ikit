package ilog

import (
	"fmt"
	"io/ioutil"
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

// SetLevel 设置日志等级，默认打印全部
func SetLevel(level Level) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if LevelDebug > level {
		debugLog.SetOutput(ioutil.Discard)
	}
	if LevelInfo > level {
		infoLog.SetOutput(ioutil.Discard)
	}
	if LevelWarn > level {
		warnLog.SetOutput(ioutil.Discard)
	}
	if LevelError > level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if LevelFatal > level {
		fatalLog.SetOutput(ioutil.Discard)
	}
	if LevelPanic > level {
		panicLog.SetOutput(ioutil.Discard)
	}
}
