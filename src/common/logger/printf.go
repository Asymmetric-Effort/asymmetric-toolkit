package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"fmt"
)

func (o *Logger) Printf(level level.LogLevel, formatString string, msg ...interface{}) {
	o.Print(level, fmt.Sprintf(formatString, msg...))
}
