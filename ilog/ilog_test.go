package ilog

import "testing"

func TestLog(t *testing.T) {
	SetLevel(LevelInfo)
	Debug()
	Info()
	Warn()
	Error()
	Fatal()
	Panic()
}
