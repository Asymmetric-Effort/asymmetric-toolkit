package logger

import (
	"fmt"
	"time"
)

func (o *Logger) ErrorInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Error,
		tags:    &tags,
		message: &message,
	})
}
