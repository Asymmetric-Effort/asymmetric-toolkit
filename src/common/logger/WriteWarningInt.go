package logger

import (
	"fmt"
	"time"
)

func (o *Logger) WarningInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Warning,
		tags:    tags,
		message: message,
	})
}
