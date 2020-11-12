package logtarget

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
