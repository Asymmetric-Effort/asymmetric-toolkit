package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)


func (o *Logger) Infof(format string, v ...interface{}){
	switch o.Level.Get() {
	case level.Info, level.Debug:
		o.Printf(level.Info, format, v...)
	case level.Critical,level.Error,level.Warning:
		break
	}
}
