package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)


func (o *Logger) Info(format string, v ...interface{}){
	if o.level.Get() >= logLevel.Info {
		o.Printf(logLevel.Info, format, v...)
	}
}
