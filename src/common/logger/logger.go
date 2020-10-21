package logger

import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

type Logger struct {
	Level    logLevel.LogLevel
	Facility LogFacility.Facility
	Writer   func(*string)
}
