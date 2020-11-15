package logger

import (
	"fmt"
	"time"
)

func (o *Logger) DebugInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Debug,
		Tags:    tags,
		Message: message,
	})
}

