package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"fmt"
	"testing"
)

func TestSpecification_AddLogLevel(t *testing.T) {

	const testDefault = logLevel.Debug

	var o Specification

	//expectedLevel := testDefault

	if logLevelArgLong != "loglevel" {
		panic("mismatch")
	}

	errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

	o.AddLogLevel(testDefault)
	errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

	errors.Assert(o.Argument[logLevelArgLong].FlagId == FlagLogLevel,
		fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
			logLevelArgLong, FlagLogLevel, o.Argument[logLevelArgLong].FlagId))

	errors.Assert(o.Argument[logLevelArgLong].Type == String,
		fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[logLevelArgLong].Type))
	errors.Assert(o.Argument[logLevelArgLong].Default == "debug",
		fmt.Sprintf("Default (debug) should be '%s' for %s",
			o.Argument[logLevelArgLong].Default,
			logLevelArgLong))
	errors.Assert(o.Argument[logLevelArgLong].Help == fmt.Sprintf(logLevelHelpText,logLevel.LevelStrings),
		"help mismatch.")

	errors.Assert(o.Argument[logLevelArgLong].Parse != nil, "Expect non-nil function pointer")

	errors.Assert(o.Argument[logLevelArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")
}
