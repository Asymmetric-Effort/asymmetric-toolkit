package logger

import (
	"fmt"
	"time"
)

func (o *Logger) Print(level LogLevel, msg string) {
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
