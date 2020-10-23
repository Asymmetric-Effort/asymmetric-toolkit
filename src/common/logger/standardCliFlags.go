package logger

import (
	logLevel "github.com/combat-diver-foundation/common/logger/log_level"
)

const (
	FlagDebug     = "debug"
	FlagDebugText = "Enable debug (verbose) logging in the logger facilities."

	FlagLogDestination        = "logdestination"
	FlagLogDestinationText    = "Specify the destination for log activity (e.g. stdout, file, syslog)"
	FlagLogDestinationDefault = "stdout"

	FlagLogFile        = "logfile"
	FlagLogFileText    = "Specify the path/filename of a log file where file is the destination."
	FlagLogFileDefault = "asymtoolkit.log"

	FlagLogLevel        = "loglevel"
	FlagLogLevelText    = "Specify the logging level (string)."
	FlagLogLevelDefault = logLevel.InfoText

	FlagLogServer        = "logserver"
	FlagLogServerText    = "Specify the FQDN and port for sending logs to a syslog server."
	FlagLogServerDefault = "udp:127.0.0.1:514"
	//ToDo: add TLS options for syslog.
)
