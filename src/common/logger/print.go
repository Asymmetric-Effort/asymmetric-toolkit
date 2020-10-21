package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"fmt"
	"time"
)

func (o *Logger) Print(level logLevel.LogLevel, msg string) {
	if o.Writer != nil {
		formattedMessage:=fmt.Sprintf(
			LogFormat,
			time.Now().String(),
			o.Facility.String(),
			level.String(),
			msg+"\n")
		o.Writer(&formattedMessage)
	}
}
