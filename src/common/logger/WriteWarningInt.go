package logger

import (
	"fmt"
	"time"
)

func (o *Logger) WarningInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Warning,
		Tags:    tags,
		Message: message,
	})
}
