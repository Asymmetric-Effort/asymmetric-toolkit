package logger

import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

type Logger struct {
	Level    level.LogLevel
	Facility LogFacility.Facility
	Writer   func(*string)
}
