package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)

func (o *Logger) Warning(format string, v ...interface{}){
	if o.level.Get() >= logLevel.Warning {
		o.Printf(logLevel.Warning, format, v...)
	}
}
