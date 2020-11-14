package logger

import (
	"fmt"
	"time"
)

func (o *Logger) WarningFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Warning,
		tags:    tags,
		message: message,
	})
}
