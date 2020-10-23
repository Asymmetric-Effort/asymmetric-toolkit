package destination

func (o *LogDestination) Get() LogDestination {
	switch *o {
	case Stdout, File, Syslog: return *o
	default:
		panic("Invalid log destination")
	}
}
