package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"fmt"
)

func (o *Logger) Printf(level logLevel.LogLevel, formatString string, msg ...interface{}) {
	o.Print(level, fmt.Sprintf(formatString, msg...))
}
