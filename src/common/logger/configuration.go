package logger

/*
	The logger facilities are configured using this common struct, which is initialized by the common/cli.
*/
import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
)

type Configuration struct {
	Level    logLevel.LogLevel
	Facility LogFacility.Facility
	Target   string
}
