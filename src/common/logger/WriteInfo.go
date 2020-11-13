package logger

func (o *Logger) Info(eventId EventId, value string, tag ...TagId){
	switch o.Level.Get() {
	case Info:
		o.Printf(Info, eventId, &tag, &value)
	}
}
