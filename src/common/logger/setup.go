package logger

func (o *Logger) Setup(config *Configuration) {
	/*
		Initialize the log destination file handle (fp)
	*/
	o.level.Set(config.Level.Get())
	o.facility.Set(config.Facility.Get())
	o.target.SetString(config.Target)
}
