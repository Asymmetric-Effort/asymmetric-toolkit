package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

func (o *Logger) Debugf(format string, v ...interface{}) {
	switch o.Level.Get() {
	case level.Debug:
		o.Printf(level.Debug, format, v...)
	case level.Critical,level.Error,level.Warning,level.Info:
		break
	}
}
