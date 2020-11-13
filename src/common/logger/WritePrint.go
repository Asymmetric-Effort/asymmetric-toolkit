package logger

import (
	"fmt"
	"time"
)

func (o *Logger) Print(level Level, msg string) {
	if o.Writer != nil {
		formattedMessage:=fmt.Sprintf(
			time.Now().String(),
			level.String(),
			msg+"\n")
		o.Writer(&formattedMessage)
	}
}
