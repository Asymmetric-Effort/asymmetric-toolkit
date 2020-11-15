package logger

import (
	"fmt"
	"time"
)

func (o *Logger) InfoInt(eventId EventId, value int, tags ...TagId) {
	message := fmt.Sprintf("%d", value)
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Info,
		Tags:    tags,
		Message: message,
	})
}
