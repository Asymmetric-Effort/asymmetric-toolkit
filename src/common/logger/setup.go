package logger

func (o *Logger) Setup(config *Configuration) {
	/*
		Initialize the log destination file handle (fp)
	*/
	o.Level.Set(config.Level.Get())
	o.Facility.Set(config.Facility.Get())
	o.Destination.Set(config.Destination)
}
