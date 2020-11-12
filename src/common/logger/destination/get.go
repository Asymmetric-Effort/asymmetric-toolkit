package logtarget

func (o *Destination) Get() Destination {
	switch *o {
	case Stdout, File, Syslog: return *o
	default:
		panic("Invalid log destination")
	}
}
