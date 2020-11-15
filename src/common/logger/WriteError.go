package logger

import "time"

/*
	Logger::Error() is the preferred logging method for error messages using only a numeric event Id and tags.
*/
func (o *Logger) Error(eventId EventId, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Error,
		Tags:    tags,
		Message: "",
	})
}
