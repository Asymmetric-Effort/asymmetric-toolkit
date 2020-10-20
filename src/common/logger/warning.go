package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

func (o *Logger) Warning(format string, v ...interface{}) {
	switch o.level.Get() {
	case logLevel.Warning, logLevel.Info, logLevel.Debug:
		o.Printf(logLevel.Warning, format, v...)
	}
}
