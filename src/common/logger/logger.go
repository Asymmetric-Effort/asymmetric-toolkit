package logger

import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
)

type Logger struct {
	level    level.Level
	facility LogFacility.Facility
	writer   func(*string)
}
