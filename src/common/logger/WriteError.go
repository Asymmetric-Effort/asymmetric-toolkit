package logger

func (o *Logger) Error(eventId EventId,value string, tag ...TagId){
	switch o.Level.Get() {
	case Debug:
		o.Printf(Error, eventId, value, tag)
	}
}
