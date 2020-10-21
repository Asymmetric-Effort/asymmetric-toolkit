package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

func (o *Logger) Warningf(format string, v ...interface{}) {
	switch o.Level.Get() {
	case level.Warning, level.Info, level.Debug:
		o.Printf(level.Warning, format, v...)
	}
}
