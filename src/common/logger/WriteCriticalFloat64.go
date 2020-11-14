package logger

import (
	"fmt"
	"time"
)

func (o *Logger) CriticalFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Critical,
		tags:    tags,
		message: message,
	})
}
