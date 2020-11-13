package logger

/*
	Commandline configuration struct.  This will be the internal state
	of the logger which will be loaded by the commandline arguments
*/
type Configuration struct {
	Level       Level              // Specifies the logging level.
	Destination Destination        // Specifies the log destination
	Settings    *map[string]string // Simple key-value store
}
