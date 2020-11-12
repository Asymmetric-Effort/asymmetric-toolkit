package logger

func (o *Logger) Error(format string, v ...interface{}) {
	switch o.Level.Get() {
	case Error, Warning, Info, Debug:
		o.Printf(Error, format, v...)
	}
}
