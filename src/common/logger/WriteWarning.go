package logger

import "time"

func (o *Logger) Warning(eventId EventId, tags ...TagId){
	o.PrintEvent(&LogEventStruct{
		EventId: eventId,
		Time:    time.Now().UnixNano(),
		Level:   Warning,
		Tags:    tags,
		Message: "",
	})
}
