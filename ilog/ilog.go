package ilog

var (
	Debug  = debugLog.Println
	Debugf = debugLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
	Warn   = warnLog.Println
	Warnf  = warnLog.Printf
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Fatal  = fatalLog.Println
	Fatalf = fatalLog.Printf
	Panic  = panicLog.Println
	Panicf = panicLog.Printf
)
