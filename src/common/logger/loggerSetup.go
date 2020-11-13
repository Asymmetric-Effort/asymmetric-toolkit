package logger
/*
	Logger::Setup() is used by an application to pass a common/logger/Configuration struct
 	containing the current configuration state of the logger into the logger for application.
	This method should validate the inputs.
 */
func (o *Logger) Setup(config *Configuration) {
	/*
		Initialize the log destination file handle (fp)
	*/
	if config == nil {
		panic("Nil logger config passed to Logger::Setup() in common/logger.")
	}
	o.Level.Set(config.Level.Get())
	o.Destination.Set(config.Destination.Get())
}
