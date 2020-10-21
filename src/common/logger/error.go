package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

func (o *Logger) Error(format string, v ...interface{}) {
	switch o.Level.Get() {
	case logLevel.Error, logLevel.Warning, logLevel.Info, logLevel.Debug:
		o.Printf(logLevel.Error, format, v...)
	}
}
