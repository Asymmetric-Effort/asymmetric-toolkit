package logger

import (
	logtarget "asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

type Logger struct {
	level    logLevel.LogLevel
	facility LogFacility.Facility
	target   logtarget.Destination
	writer   func(*string)
}
