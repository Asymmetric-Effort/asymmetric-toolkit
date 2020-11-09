package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddSourcePattern(t *testing.T) {

	const testDefault = ".+"

	func() {
		var o Specification

		if sourcePatternArgLong != "sourcePattern" {
			panic("mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddSourcePattern(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[sourcePatternArgLong].FlagId == flagSourcePattern,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
				sourcePatternArgLong, flagSourcePattern, o.Argument[sourcePatternArgLong].FlagId))

		errors.Assert(o.Argument[sourcePatternArgLong].Type == String,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[sourcePatternArgLong].Type))

		errors.Assert(o.Argument[sourcePatternArgLong].Default == testDefault,
			fmt.Sprintf("Default should be empty for %s", sourcePatternArgLong))

		errors.Assert(o.Argument[sourcePatternArgLong].Help == sourcePatternHelpText,
			"Declares a source mode (random, sequential, dictionary).")

		errors.Assert(o.Argument[sourcePatternArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[sourcePatternArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

		if o.Argument[sourcePatternArgLong].Default != testDefault {
			panic(fmt.Sprintf("Default should be false for domain flag (%d)", flagSourcePattern))
		}
	}()
}
