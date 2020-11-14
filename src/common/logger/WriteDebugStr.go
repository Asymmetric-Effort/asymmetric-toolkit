package logger

import "time"

func (o *Logger) DebugStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Debug,
		tags:    tags,
		message: message,
	})
}
