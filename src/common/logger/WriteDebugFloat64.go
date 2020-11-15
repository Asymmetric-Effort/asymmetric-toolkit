package logger

import (
	"fmt"
	"time"
)

func (o *Logger) DebugFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Debug,
		Tags:    tags,
		Message: message,
	})
}

