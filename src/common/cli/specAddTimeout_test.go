package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddTimeout(t *testing.T) {

	const testTimeoutDefault = 10

	func() {
		var o Specification

		if timeoutArgLong != "timeout" {
			panic("timeoutArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddTimeout(testTimeoutDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[timeoutArgLong].FlagId == FlagTimeout,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", timeoutArgLong, FlagTimeout,
				o.Argument[timeoutArgLong].FlagId))

		errors.Assert(o.Argument[timeoutArgLong].Type == Integer,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[timeoutArgLong].Type))

		errors.Assert(o.Argument[timeoutArgLong].Default == strconv.Itoa(testTimeoutDefault),
			fmt.Sprintf("Default should be empty for %s", timeoutArgLong))

		errors.Assert(o.Argument[timeoutArgLong].Help == timeoutHelpText,
			"Indicates the connection timeout (in seconds)")

		errors.Assert(o.Argument[timeoutArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[timeoutArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		val, err := strconv.Atoi(o.Argument[timeoutArgLong].Default)
		if err != nil {
			panic(err)
		}
		if val != testTimeoutDefault {
			panic(fmt.Sprintf("Default should be false for timeout flag (%d): '%d'", FlagTimeout, val))
		}
	}()
}
