package logger
/*
	Destination::Set() is used to set the numeric Destination value.
 */
func (o *Destination) Set(d Destination) {
	switch Destination(d) {
	case Stdout, File, Syslog:
		*o = Destination(d)
	default:
		panic("Invalid log destination")
	}
}