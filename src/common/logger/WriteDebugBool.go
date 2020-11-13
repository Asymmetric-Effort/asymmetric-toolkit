package logger

import (
	"fmt"
	"time"
)

func (o *Logger) DebugBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Debug,
		tags:    &tags,
		message: &message,
	})
}
