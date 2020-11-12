package logtarget

func (o *Destination) Set(d Destination) {
	switch Destination(d) {
	case Stdout, File, Syslog:
		*o = Destination(d)
	default:
		panic("Invalid log destination")
	}
}
