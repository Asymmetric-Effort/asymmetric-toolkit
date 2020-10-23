package logger

import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/tools/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
)

type Logger struct {
	level    logLevel.LogLevel
	facility LogFacility.Facility
	writer func(*string)
}
