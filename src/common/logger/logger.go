package logger
/*
	Logger is a struct at the highest level of the log management facilities.
	This structure contains the configured state of the application for logging, as
	established by the Logger::Setup() method.
 */

type Logger struct {
	Level       LogLevel      // The current Log level (Critical, Error, Warning, Info, Debug) as a numeric value.
	Facility    Facility      // A numeric identifier indicating the name of a given logger instance or code area.
	Destination Destination   // The numeric identifier of the log target (protocol) where logs will ship.
	Settings    *LogSettings  // arbitrary log settings map pointer when a logging subsystem requires more parameters.
	Writer      func(*string) // Payload log writer function (configured as per Destination.

	// Private properties
	nextTag TagId //This is the next tag to be issued by Logger::TagCreate()
}
