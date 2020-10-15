package destination

func (o *LogDestination) Set(d LogDestination){
	switch LogDestination(d) {
	case Stdout, File, Syslog: *o = LogDestination(d)
	default:
		panic("Invalid log destination")
	}
}
