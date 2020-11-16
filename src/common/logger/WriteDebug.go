package logger

import "time"

/*
	Logger::Debug() is the preferred logging method for debug messages using only a numeric event Id and tags.
*/
func (o *Logger) Debug(eventId EventId, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Debug,
		Tags:    tags,
		Message: "",
	})
}