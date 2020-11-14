package logger

import "time"

func (o *Logger) InfoStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level: Info,
		tags:    tags,
		message: message,
	})
}
