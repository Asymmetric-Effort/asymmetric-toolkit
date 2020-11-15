package logger

import (
	"fmt"
	"time"
)

func (o *Logger) CriticalFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Critical,
		Tags:    tags,
		Message: message,
	})
}
