package logger

import "time"

func (o *Logger) InfoStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level: Info,
		Tags:    tags,
		Message: message,
	})
}
