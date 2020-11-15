package logger

import (
	"fmt"
	"time"
)

func (o *Logger) WarningFloat64(eventId EventId, value float64, tags ...TagId) {
	message := fmt.Sprintf("%f", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Warning,
		Tags:    tags,
		Message: message,
	})
}
