package logger

func (o *Logger) Error(eventId EventId, value string, tag ...TagId){
	switch o.Level.Get() {
	case Error:
		o.Printf(Error, eventId, &tag, &value)
	}
}
