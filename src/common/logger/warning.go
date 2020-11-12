package logger

func (o *Logger) Warning(format string, v ...interface{}) {
	switch o.Level.Get() {
	case Warning, Info, Debug:
		o.Printf(Warning, format, v...)
	}
}
