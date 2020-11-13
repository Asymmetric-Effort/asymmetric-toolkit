package logger

import (
	"fmt"
	"time"
)

func (o *Logger) ErrorFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Error,
		tags:    &tags,
		message: &message,
	})
}
