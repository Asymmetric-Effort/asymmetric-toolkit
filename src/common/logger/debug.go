package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

func (o *Logger) Debug(format string, v ...interface{}){
	switch o.Level.Get() {
	case logLevel.Debug:
		o.Printf(logLevel.Debug, format, v...)
	}
}
