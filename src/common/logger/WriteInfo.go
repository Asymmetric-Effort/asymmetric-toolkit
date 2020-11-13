package logger

func (o *Logger) Info(eventId EventId,value string, tag ...TagId){
	switch o.Level.Get() {
	case Debug:
		o.Printf(Info, eventId, value, tag)
	}
}
