package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"fmt"
	"time"
)

func (o *Logger) Print(level level.Level, msg string) {
	if o.writer != nil {
		formattedMessage:=fmt.Sprintf(
			LogFormat,
			time.Now().String(),
			o.facility.String(),
			level.String(),
			msg+"\n")
		o.writer(&formattedMessage)
	}
}
