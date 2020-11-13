package logger

import "time"

/*
	Logger::Info() is the preferred logging method for informational messages using only a numeric event Id and tags.
*/
func (o *Logger) Info(eventId EventId, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Info,
		tags:    &tags,
		message: nil,
	})
}
