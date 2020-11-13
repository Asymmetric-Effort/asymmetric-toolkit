package logger

import (
	"fmt"
	"time"
)

func (o *Logger) InfoInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Info,
		tags:    &tags,
		message: &message,
	})
}
