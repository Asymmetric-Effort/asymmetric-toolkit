package logger

import (
	"fmt"
	"time"
)

func (o *Logger) CriticalBool(eventId EventId, value bool, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Critical,
		Tags:    tags,
		Message: fmt.Sprintf("%v", value),
	})
}
