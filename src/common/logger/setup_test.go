package logger_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"testing"
)

func TestLoggerSetupStdOutHappy(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(logLevel.Debug)
	log.Setup(&config)
	errors.Assert(log.Facility.String() == logger.DefaultLoggerFacility, "Unexpected initial logger facility.")
	errors.Assert(log.Level.Get() == logLevel.Debug, "Expected DEBUG level.")
}
func TestLoggerSetupStdOutNoDestination(t *testing.T) {
	//Default to stdout
	var log logger.Logger
	var config cli.Configuration
	config.Log.Level.Set(logLevel.Debug)
	log.Setup(&config)
	errors.Assert(log.Facility.String() == logger.DefaultLoggerFacility, "Unexpected initial logger facility.")
	errors.Assert(log.Level.Get() == logLevel.Debug, "Expected DEBUG level.")
}

func TestLoggerSetupStdOutBadDestination(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	const badDestination = -1
	defer func(){recover()}()
	config.Log.Destination = badDestination
	config.Log.Level.Set(logLevel.Debug)
	log.Setup(&config)
	errors.Assert(log.Facility.String() == logger.DefaultLoggerFacility, "Unexpected initial logger facility.")
	errors.Assert(log.Level.Get() == logLevel.Debug, "Expected DEBUG level.")
}
