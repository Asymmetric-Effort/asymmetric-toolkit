package logger

import "time"

func (o *Logger) WarningStr(eventId EventId, message string, tags ...TagId) {
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Warning,
		Tags:    tags,
		Message: message,
	})
}
