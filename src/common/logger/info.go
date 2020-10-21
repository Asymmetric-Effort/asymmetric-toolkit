package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)


func (o *Logger) Info(format string, v ...interface{}){
	switch o.Level.Get() {
	case logLevel.Info, logLevel.Debug:
		o.Printf(logLevel.Info, format, v...)
	}
}
