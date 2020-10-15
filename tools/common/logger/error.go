package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)

func (o *Logger) Error(format string, v ...interface{}){
	if o.level.Get() >= logLevel.Error {
		o.Printf(logLevel.Error, format, v...)
	}
}
