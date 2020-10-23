package LogDestination

type Target int

const (
	Stdout Target = 0
	File   Target = 1
	Syslog Target = 2

	StdoutText string = "stdout"
	FileText   string = "file"
	SyslogText string = "syslog"
)
