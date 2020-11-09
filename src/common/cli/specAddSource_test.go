package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddSource(t *testing.T) {

	const testDefault = "sequence"

	func() {
		var o Specification

		if sourceArgLong != "source" {
			panic("mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddSource(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[sourceArgLong].FlagId == flagSource,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
				sourceArgLong, flagSource, o.Argument[sourceArgLong].FlagId))

		errors.Assert(o.Argument[sourceArgLong].Type == String,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[sourceArgLong].Type))

		errors.Assert(o.Argument[sourceArgLong].Default == testDefault,
			fmt.Sprintf("Default should be empty for %s", sourceArgLong))

		errors.Assert(o.Argument[sourceArgLong].Help == sourceHelpText,
			"Declares a source mode (random, sequential, dictionary).")

		errors.Assert(o.Argument[sourceArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[sourceArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

		if o.Argument[sourceArgLong].Default != testDefault {
			panic(fmt.Sprintf("Default should be false for domain flag (%d)", flagSource))
		}
	}()
}
