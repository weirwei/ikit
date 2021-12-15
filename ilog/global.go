package ilog

type Level int

// 日志等级
const (
	LevelPanic Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// nolint
// 超级颜色
const (
	none        = "\033[0m"
	black       = "\033[0;30m"
	darkGray    = "\033[1;30m"
	blue        = "\033[0;34m"
	lightBlue   = "\033[1;34m"
	green       = "\033[0;32m"
	lightGreen  = "\033[1;32m"
	cyan        = "\033[0;36m"
	lightCyan   = "\033[1;36m"
	red         = "\033[0;31m"
	lightRed    = "\033[1;31m"
	purple      = "\033[0;35m"
	lightPurple = "\033[1;35m"
	brown       = "\033[0;33m"
	yellow      = "\033[1;33m"
	lightGray   = "\033[0;37m"
	white       = "\033[1;37m"
)
