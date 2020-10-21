package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

func (o *Logger) Debug(format string, v ...interface{}){
	switch o.level.Get() {
	case level.Debug:
		o.Printf(level.Debug, format, v...)
	}
}
