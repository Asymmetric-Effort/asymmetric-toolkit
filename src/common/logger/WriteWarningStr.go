package logger

import "time"

func (o *Logger) WarningStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Warning,
		tags:    tags,
		message: message,
	})
}
