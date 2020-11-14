package logger

import (
	"fmt"
	"time"
)

func (o *Logger) CriticalBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Critical,
		tags:    tags,
		message: message,
	})
}
