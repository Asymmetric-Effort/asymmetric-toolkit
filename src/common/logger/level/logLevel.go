package level

type Level int

const (
	Critical Level = 0
	Error    Level = 1
	Warning  Level = 2
	Info     Level = 3
	Debug    Level = 4

	Strings string = "CRIT,ERROR,WARN,INFO,DEBUG"
)
