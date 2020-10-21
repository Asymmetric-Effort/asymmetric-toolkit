package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

func (o *Logger) Errorf(format string, v ...interface{}) {
	switch o.Level.Get() {
	case level.Error, level.Warning, level.Info, level.Debug:
		o.Printf(level.Error, format, v...)
	case level.Critical:
		break
	}
}
