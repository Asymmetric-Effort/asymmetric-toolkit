package logger

func (o *Logger) Warning(eventId EventId,value string, tag ...TagId){
	switch o.Level.Get() {
	case Debug:
		o.Printf(Warning, eventId, value, tag)
	}
}
