package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddOutput(t *testing.T) {

	const testDefault = "localdomain"

	func() {
		var o Specification

		if outputArgLong != "output" {
			panic("mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddOutput(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[outputArgLong].FlagId == FlagOutput,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
				outputArgLong, FlagOutput, o.Argument[outputArgLong].FlagId))

		errors.Assert(o.Argument[outputArgLong].Type == String,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[outputArgLong].Type))

		errors.Assert(o.Argument[outputArgLong].Default == testDefault,
			fmt.Sprintf("Default should be empty for %s", outputArgLong))

		errors.Assert(o.Argument[outputArgLong].Help == outputHelpText,  "Specifies a fully qualified domain name.")

		errors.Assert(o.Argument[outputArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[outputArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		if o.Argument[outputArgLong].Default != testDefault {
			panic(fmt.Sprintf("Default should be false for output (%d)", FlagOutput))
		}
	}()
}
