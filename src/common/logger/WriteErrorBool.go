package logger

import (
	"fmt"
	"time"
)

func (o *Logger) ErrorBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Error,
		tags:    &tags,
		message: &message,
	})
}
