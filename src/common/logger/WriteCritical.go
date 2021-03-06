package logger

/*
	Logger::Critical() is the preferred logging method for critical messages using only a numeric event Id and tags.
*/

import "time"

func (o *Logger) Critical(eventId EventId, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Critical,
		Tags:    tags,
		Message: "",
	})
}
