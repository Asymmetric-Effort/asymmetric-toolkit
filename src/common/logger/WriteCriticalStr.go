package logger

import "time"

func (o *Logger) CriticalStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Critical,
		tags:    tags,
		message: message,
	})
}
