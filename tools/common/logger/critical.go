package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)

func (o *Logger) Critical(format string, v ...interface{}){
	if o.level.Get() >= logLevel.Critical {
		o.Printf(logLevel.Critical, format, v...)
	}
}
