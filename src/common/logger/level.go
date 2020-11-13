package logger
/*
	Level is a numeric type showing log level.
 */

type Level int

const (
	Critical Level = 0
	Error    Level = 1
	Warning  Level = 2
	Info     Level = 3
	Debug    Level = 4

	LevelStrings string = "critical,error,warning,info,debug"
)

