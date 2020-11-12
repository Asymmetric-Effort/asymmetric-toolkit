package logger

func (o *Logger) Info(format string, v ...interface{}){
	switch o.Level.Get() {
	case Info, Debug:
		o.Printf(Info, format, v...)
	}
}
