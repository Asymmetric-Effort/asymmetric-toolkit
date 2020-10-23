package destination

type LogDestination int

const (
	Stdout LogDestination = 0
	File   LogDestination = 1
	Syslog LogDestination = 2
)
