package logger

import "time"

/*
	Logger::Info() is the preferred logging method for informational messages using only a numeric event Id and tags.
*/
func (o *Logger) Info(eventId EventId, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Info,
		Tags:    tags,
		Message: "",
	})
}
