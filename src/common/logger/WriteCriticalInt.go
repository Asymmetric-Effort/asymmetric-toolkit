package logger

import (
	"fmt"
	"time"
)

func (o *Logger) CriticalInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Critical,
		Tags:    tags,
		Message: message,
	})
}
