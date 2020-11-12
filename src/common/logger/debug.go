package logger

func (o *Logger) Debug(format string, v ...interface{}){
	switch o.Level.Get() {
	case Debug:
		o.Printf(Debug, format, v...)
	}
}
