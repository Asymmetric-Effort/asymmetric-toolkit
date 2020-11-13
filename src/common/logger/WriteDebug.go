package logger

func (o *Logger) Debug(eventId EventId,value string, tag ...TagId){
	switch o.Level.Get() {
	case Debug:
		o.Printf(Debug, eventId, value, tag)
	}
}
