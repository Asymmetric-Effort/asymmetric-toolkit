package logger

/*
	Commandline configuration struct.  This will be the internal state
	of the logger which will be loaded by the commandline arguments
*/
type Configuration struct {
	Level       LogLevel
	Facility    Facility
	Destination Destination
}
