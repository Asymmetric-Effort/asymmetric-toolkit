package logger

import (
	"fmt"
	"time"
)

func (o *Logger) DebugInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Debug,
		tags:    tags,
		message: message,
	})
}

