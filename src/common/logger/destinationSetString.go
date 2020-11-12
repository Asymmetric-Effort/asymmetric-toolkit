package logger

/*
	Destination::SetString() will set the internal logging destination based on an input string value.
*/

import "strings"

func (o *Destination) SetString(d string) {
	switch strings.ToLower(strings.TrimSpace(d)) {
	case "stdout":
		*o = Stdout
	case "file":
		*o = File
	case "syslog":
		*o = Syslog
	default:
		panic("Invalid log destination")
	}
}
