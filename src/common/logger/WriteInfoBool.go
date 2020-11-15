package logger

import (
	"fmt"
	"time"
)

func (o *Logger) InfoBool(eventId EventId, value bool, tags ...TagId) {
	message := fmt.Sprintf("%t", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Info,
		Tags:    tags,
		Message: message,
	})
}
