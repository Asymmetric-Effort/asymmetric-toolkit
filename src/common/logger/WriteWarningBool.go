package logger

import (
	"fmt"
	"time"
)

func (o *Logger) WarningBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Warning,
		Tags:    tags,
		Message: message,
	})
}
