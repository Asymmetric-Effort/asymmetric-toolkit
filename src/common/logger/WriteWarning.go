package logger

import "time"

func (o *Logger) Warning(eventId EventId, value string, tags ...TagId){
	o.PrintEvent(&LogEventStruct{
		eventId: eventId,
		time:    time.Now(),
		level:   Warning,
		tags:    tags,
		message: "",
	})
}
