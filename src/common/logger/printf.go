package logger

import (
	"fmt"
)

func (o *Logger) Printf(level LogLevel, formatString string, msg ...interface{}) {
	o.Print(level, fmt.Sprintf(formatString, msg...))
}
