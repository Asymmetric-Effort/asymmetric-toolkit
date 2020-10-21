package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

func (o *Logger) Warning(format string, v ...interface{}) {
	switch o.level.Get() {
	case level.Warning, level.Info, level.Debug:
		o.Printf(level.Warning, format, v...)
	}
}
