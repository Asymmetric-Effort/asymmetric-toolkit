package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
	"fmt"
	"time"
)

func (o *Logger) Print(level logLevel.LogLevel, msg string) {
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
