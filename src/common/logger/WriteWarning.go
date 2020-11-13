package logger

func (o *Logger) Warning(eventId EventId, value string, tag ...TagId){
	switch o.Level.Get() {
	case Warning:
		o.Printf(Warning, eventId, &tag, &value)
	}
}
