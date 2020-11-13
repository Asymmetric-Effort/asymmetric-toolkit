package logger

func (o *Logger) Critical(eventId EventId,value string, tag ...TagId){
	switch o.Level.Get() {
	case Debug:
		o.Printf(Critical, eventId, value, tag)
	}
}
