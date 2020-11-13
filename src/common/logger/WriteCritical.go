package logger

func (o *Logger) Critical(eventId EventId, value string, tag ...TagId){
	switch o.Level.Get() {
	case Critical:
		o.Printf(Critical, eventId, &tag, &value)
	}
}
