package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"testing"
)

func TestProcessSpecification_Configuration(t *testing.T) {
	var config Configuration
	errors.Assert(!config.Force, "expected config.Force false")
	errors.Assert(!config.Debug, "expected config.Debug false")
	errors.Assert(config.InputFile == "", "expected empty string")
	errors.Assert(config.OutputFile == "", "expected empty string")
	var c logger.Configuration
	errors.Assert(config.Log == c, "Expected default log config")
}
