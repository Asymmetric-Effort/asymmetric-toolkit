package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddMaxWordSize(t *testing.T) {

	const testDefault = 10

	func() {
		var o Specification

		if maxWordSizeArgLong != "maxWordSize" {
			panic("maxWordSizeArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddMaxWordSize(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[maxWordSizeArgLong].FlagId == FlagSourceMaxWordSize,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", maxWordSizeArgLong, FlagSourceMaxWordSize,
				o.Argument[maxWordSizeArgLong].FlagId))

		errors.Assert(o.Argument[maxWordSizeArgLong].Type == Integer,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[maxWordSizeArgLong].Type))

		errors.Assert(o.Argument[maxWordSizeArgLong].Default == strconv.Itoa(testDefault),
			fmt.Sprintf("Default should be empty for %s", maxWordSizeArgLong))

		errors.Assert(o.Argument[maxWordSizeArgLong].Help == maxWordSizeHelpText,
			"help text mismatch")

		errors.Assert(o.Argument[maxWordSizeArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[maxWordSizeArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		val, err := strconv.Atoi(o.Argument[maxWordSizeArgLong].Default)
		if err != nil {
			panic(err)
		}
		if val != testDefault {
			panic(fmt.Sprintf("Default should be false for delay Flag (%d): '%d'", FlagSourceMaxWordSize, val))
		}
	}()
}
