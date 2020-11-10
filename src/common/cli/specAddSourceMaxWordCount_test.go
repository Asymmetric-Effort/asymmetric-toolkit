package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddMaxWordCount(t *testing.T) {

	const testDefault = 10

	func() {
		var o Specification

		if maxWordCountArgLong != "maxWordCount" {
			panic("maxWordCountArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddMaxWordCount(testDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[maxWordCountArgLong].FlagId == FlagSourceMaxWordCount,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", maxWordCountArgLong, FlagSourceMaxWordCount,
				o.Argument[maxWordCountArgLong].FlagId))

		errors.Assert(o.Argument[maxWordCountArgLong].Type == Integer,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[maxWordCountArgLong].Type))

		errors.Assert(o.Argument[maxWordCountArgLong].Default == strconv.Itoa(testDefault),
			fmt.Sprintf("Default should be empty for %s", maxWordCountArgLong))

		errors.Assert(o.Argument[maxWordCountArgLong].Help == maxWordCountHelpText,
			"Indicates the delay each attacker will observe between attacks")

		errors.Assert(o.Argument[maxWordCountArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[maxWordCountArgLong].Expects == ExpectNone, "Next expected should be ExpectNone")

		val, err := strconv.Atoi(o.Argument[maxWordCountArgLong].Default)
		if err != nil {
			panic(err)
		}
		if val != testDefault {
			panic(fmt.Sprintf("Default should be false for delay flag (%d): '%d'", FlagSourceMaxWordCount, val))
		}
	}()
}
