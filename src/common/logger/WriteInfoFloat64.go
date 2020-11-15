package logger

import (
	"fmt"
	"time"
)

func (o *Logger) InfoFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Info,
		Tags:    tags,
		Message: message,
	})
}
