package logger

import (
	"fmt"
	"time"
)

func (o *Logger) WarningBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Warning,
		tags:    tags,
		message: message,
	})
}
