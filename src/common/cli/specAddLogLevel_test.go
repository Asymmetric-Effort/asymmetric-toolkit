package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddLogLevel(t *testing.T) {

	const testDefault = "localdomain"

	func() {
		var o Specification

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

		errors.Assert(o.Argument[logLevelArgLong].Default == testDefault,
			fmt.Sprintf("Default should be empty for %s", logLevelArgLong))

		errors.Assert(o.Argument[logLevelArgLong].Help == logLevelHelpText, "help mismatch.")

		errors.Assert(o.Argument[logLevelArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[logLevelArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		if o.Argument[logLevelArgLong].Default != testDefault {
			panic(fmt.Sprintf("Default should be false for log level flag (%d)", FlagLogLevel))
		}
	}()
}
