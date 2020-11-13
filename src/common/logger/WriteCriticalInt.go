package logger

import (
	"fmt"
	"time"
)

func (o *Logger) CriticalInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Critical,
		tags:    &tags,
		message: &message,
	})
}
