package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddMinWordSize(t *testing.T) {

	const testDefault = 10

	func() {
		var o Specification

		if minWordSizeArgLong != "minWordSize" {
			panic("minWordSizeArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddMinWordSize(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[minWordSizeArgLong].FlagId == flagSourceMinWordSize,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", minWordSizeArgLong, flagSourceMinWordSize,
				o.Argument[minWordSizeArgLong].FlagId))

		errors.Assert(o.Argument[minWordSizeArgLong].Type == Integer,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[minWordSizeArgLong].Type))

		errors.Assert(o.Argument[minWordSizeArgLong].Default == strconv.Itoa(testDefault),
			fmt.Sprintf("Default should be empty for %s", minWordSizeArgLong))

		errors.Assert(o.Argument[minWordSizeArgLong].Help == minWordSizeHelpText,
			"help text mismatch")

		errors.Assert(o.Argument[minWordSizeArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[minWordSizeArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

		val, err := strconv.Atoi(o.Argument[minWordSizeArgLong].Default)
		if err != nil {
			panic(err)
		}
		if val != testDefault {
			panic(fmt.Sprintf("Default should be false for delay flag (%d): '%d'", flagSourceMinWordSize, val))
		}
	}()
}
