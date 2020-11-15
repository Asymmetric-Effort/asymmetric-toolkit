package logger

import (
	"fmt"
	"time"
)

func (o *Logger) ErrorFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Error,
		Tags:    tags,
		Message: message,
	})
}
