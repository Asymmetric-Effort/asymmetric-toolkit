package logLevel

type LogLevel int

const (
	Critical LogLevel = 0
	Error    LogLevel = 1
	Warning  LogLevel = 2
	Info     LogLevel = 3
	Debug    LogLevel = 4

	LevelStrings string = "critical,error,warning,info,debug"
)