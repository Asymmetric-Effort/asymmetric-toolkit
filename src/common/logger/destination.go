package logger

/*
	Destination struct defines the data type and constant values for logging destinations.
	ToDo: we need to add configuration parameters for the destination when log files and syslog.
*/

type Destination int

const (
	Stdout Destination = 0
	File   Destination = 1
	Syslog Destination = 2
)
