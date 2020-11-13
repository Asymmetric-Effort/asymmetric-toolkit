package logger

import "time"

/*
	Logger::Debug() is the preferred logging method for debug messages using only a numeric event Id and tags.
*/
func (o *Logger) Debug(eventId EventId, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Debug,
		tags:    &tags,
		message: nil,
	})
}
