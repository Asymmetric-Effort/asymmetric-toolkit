package logger
/*
	Destination::String() return the string value of the current logger destination.
 */
func (o *Destination) String() string {
	switch *o {
	case Stdout: return "stdout"
	case File: return "file"
	case Syslog: return "syslog"
	default:
		panic("Invalid destination")
	}
}
