package logger

import "time"

func (o *Logger) DebugStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Debug,
		Tags:    tags,
		Message: message,
	})
}
