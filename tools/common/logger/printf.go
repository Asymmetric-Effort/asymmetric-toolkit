package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
	"fmt"
)

func (o *Logger) Printf(level logLevel.LogLevel, formatString string, msg ...interface{}) {
	o.Print(level, fmt.Sprintf(formatString, msg...))
}
