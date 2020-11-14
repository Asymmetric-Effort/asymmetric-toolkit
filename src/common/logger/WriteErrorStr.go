package logger

import "time"

func (o *Logger) ErrorStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Error,
		tags:    tags,
		message: message,
	})
}
