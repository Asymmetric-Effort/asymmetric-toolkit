package logger

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger/logLevel"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"testing"
)

func TestLoggerSetupStdOutHappy(t *testing.T) {
	var log Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(logLevel.Debug)
	log.Setup(&config)
	errors.Assert(log.facility.String() == defaultLoggerFacility, "Unexpected initial logger facility.")
	errors.Assert(log.level.Get() == logLevel.Debug, "Expected DEBUG level.")
}
func TestLoggerSetupStdOutNoDestination(t *testing.T) {
	//Default to stdout
	var log Logger
	var config cli.Configuration
	config.Log.Level.Set(logLevel.Debug)
	log.Setup(&config)
	errors.Assert(log.facility.String() == defaultLoggerFacility, "Unexpected initial logger facility.")
	errors.Assert(log.level.Get() == logLevel.Debug, "Expected DEBUG level.")
}

func TestLoggerSetupStdOutBadDestination(t *testing.T) {
	var log Logger
	var config cli.Configuration
	const badDestination = -1
	defer func(){recover()}()
	config.Log.Destination = badDestination
	config.Log.Level.Set(logLevel.Debug)
	log.Setup(&config)
	errors.Assert(log.facility.String() == defaultLoggerFacility, "Unexpected initial logger facility.")
	errors.Assert(log.level.Get() == logLevel.Debug, "Expected DEBUG level.")
}
