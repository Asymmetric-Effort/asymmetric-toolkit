package logtarget

type Destination int

const (
	Stdout Destination = 0
	File   Destination = 1
	Syslog Destination = 2
)
