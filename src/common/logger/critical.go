package logger

func (o *Logger) Critical(format string, v ...interface{}) {
	switch o.Level.Get() {
	case Critical, Error, Warning, Info, Debug:
		o.Printf(Critical, format, v...)
	}
}
