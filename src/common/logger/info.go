package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)


func (o *Logger) Info(format string, v ...interface{}){
	switch o.level.Get() {
	case level.Info, level.Debug:
		o.Printf(level.Info, format, v...)
	}
}
