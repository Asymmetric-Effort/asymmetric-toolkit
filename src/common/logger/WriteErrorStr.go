package logger

import "time"

func (o *Logger) ErrorStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Error,
		Tags:    tags,
		Message: message,
	})
}
