package logger

import (
	"fmt"
	"time"
)

func (o *Logger) InfoFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Info,
		tags:    &tags,
		message: &message,
	})
}
