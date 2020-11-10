package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddDelay(t *testing.T) {

	const testDelayDefault = 10

	func() {
		var o Specification

		if delayArgLong != "delay" {
			panic("delayArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddDelay(testDelayDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[delayArgLong].FlagId == FlagDelay,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", delayArgLong, FlagDelay,
				o.Argument[delayArgLong].FlagId))

		errors.Assert(o.Argument[delayArgLong].Type == Integer,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[delayArgLong].Type))

		errors.Assert(o.Argument[delayArgLong].Default == strconv.Itoa(testDelayDefault),
			fmt.Sprintf("Default should be empty for %s", delayArgLong))

		errors.Assert(o.Argument[delayArgLong].Help == delayHelpText,
			"Indicates the delay each attacker will observe between attacks")

		errors.Assert(o.Argument[delayArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[delayArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		val, err := strconv.Atoi(o.Argument[delayArgLong].Default)
		if err != nil {
			panic(err)
		}
		if val != testDelayDefault {
			panic(fmt.Sprintf("Default should be false for delay flag (%d): '%d'", FlagDelay, val))
		}
	}()
}
