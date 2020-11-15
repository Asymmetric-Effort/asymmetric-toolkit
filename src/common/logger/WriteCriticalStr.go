package logger

import "time"

func (o *Logger) CriticalStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Critical,
		Tags:    tags,
		Message: message,
	})
}
