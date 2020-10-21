package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

func (o *Logger) Critical(format string, v ...interface{}) {
	switch o.Level.Get() {
	case level.Critical, level.Error, level.Warning, level.Info, level.Debug:
		o.Printf(level.Critical, format, v...)
	}
}
