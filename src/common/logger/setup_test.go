package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLoggerSetupStdOutHappy(t *testing.T) {
	var log Logger
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)
	log.Setup(&config)
	errors.Assert(log.Facility.String() == defaultLoggerFacility, "Unexpected initial logger Facility.")
	errors.Assert(log.Level.Get() == Debug, "Expected DEBUG Level.")
}
func TestLoggerSetupStdOutNoDestination(t *testing.T) {
	//Default to stdout
	var log Logger
	var config Configuration
	config.Level.Set(Debug)
	log.Setup(&config)
	errors.Assert(log.Facility.String() == defaultLoggerFacility, "Unexpected initial logger Facility.")
	errors.Assert(log.Level.Get() == Debug, "Expected DEBUG Level.")
}

func TestLoggerSetupStdOutBadDestination(t *testing.T) {
	var log Logger
	var config Configuration
	const badDestination = -1
	defer func(){recover()}()
	config.Destination = badDestination
	config.Level.Set(Debug)
	log.Setup(&config)
	errors.Assert(log.Facility.String() == defaultLoggerFacility, "Unexpected initial logger Facility.")
	errors.Assert(log.Level.Get() == Debug, "Expected DEBUG Level.")
}
