package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)

func (o *Logger) Debug(format string, v ...interface{}){
	if o.level.Get() >= logLevel.Debug {
		o.Printf(logLevel.Debug, format, v...)
	}
}
