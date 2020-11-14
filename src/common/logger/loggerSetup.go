package logger

/*
	Logger::Setup() is used by an application to pass a common/logger/Configuration struct
 	containing the current configuration state of the logger into the logger for application.
	This method should validate the inputs.
*/
func (o *Logger) Setup(config *Configuration) { //given a logger configuration...
	if config == nil {
		panic("Nil logger config passed to Logger::Setup() in common/logger.")
	}
	//
	// Configure the logging level.
	//
	o.Level.Set(config.Level.Get())
	//
	// Evaluate the log destination and connect to our Writer function.
	//
	switch config.Destination.Get() {
	case Stdout:
		o.Writer = o.logWriterStdOut
	case File:
		panic("Log Writer File not implemented")
		//o.Writer = o.logWriterFile
	case Syslog:
		panic("Log Writer Syslog not implemented")
		//o.Writer = o.logWriterSyslog
	default:
		panic("Invalid logging destination.")
	}
	o.tags = TagTracker{
		nextTag:0,
		global:TagTable{},
		tagNames:make(TagNameDictionary,1),
		tagIds:make(TagDictionary,1),
	}
}
