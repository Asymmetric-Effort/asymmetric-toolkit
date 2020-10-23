package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

func (o *Logger) Critical(format string, v ...interface{}) {
	switch o.level.Get() {
	case logLevel.Critical, logLevel.Error, logLevel.Warning, logLevel.Info, logLevel.Debug:
		o.Printf(logLevel.Critical, format, v...)
	}
}
