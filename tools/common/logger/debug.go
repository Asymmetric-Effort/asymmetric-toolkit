package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)

func (o *Logger) Debug(format string, v ...interface{}){
	switch o.level.Get() {
	case logLevel.Debug:
		o.Printf(logLevel.Debug, format, v...)
	}
}
