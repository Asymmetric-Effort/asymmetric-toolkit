package logger

import (
	"fmt"
	"time"
)

func (o *Logger) DebugFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Debug,
		tags:    tags,
		message: message,
	})
}

