package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_AddSourceDictionary(t *testing.T) {

	const testDefault = "my_dictionary.dat"

	func() {
		var o Specification

		if sourceDictionaryArgLong != "dictionary" {
			panic("mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddSourceDictionary(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[sourceDictionaryArgLong].FlagId == FlagSourceDictionary,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d",
				sourceDictionaryArgLong, FlagSourceDictionary, o.Argument[sourceDictionaryArgLong].FlagId))

		errors.Assert(o.Argument[sourceDictionaryArgLong].Type == String,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[sourceDictionaryArgLong].Type))

		errors.Assert(o.Argument[sourceDictionaryArgLong].Default == testDefault,
			fmt.Sprintf("Default should be empty for %s", sourceDictionaryArgLong))

		errors.Assert(o.Argument[sourceDictionaryArgLong].Help == sourceDictionaryHelpText,
			"Declares a source mode (random, sequential, dictionary)")

		errors.Assert(o.Argument[sourceDictionaryArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[sourceDictionaryArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

		if o.Argument[sourceDictionaryArgLong].Default != testDefault {
			panic(fmt.Sprintf("Default should be false for domain flag (%d)", FlagSourceDictionary))
		}
	}()
}
