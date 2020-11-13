package logger

import (
	"fmt"
	"time"
)

func (o *Logger) InfoBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Info,
		tags:    &tags,
		message: &message,
	})
}
