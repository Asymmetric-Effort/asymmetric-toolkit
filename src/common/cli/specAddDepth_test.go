package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestSpecification_AddDepth(t *testing.T) {

	const testDepthDefault = 10

	func() {
		var o Specification

		if depthArgLong != "depth" {
			panic("depthArgLong mismatch")
		}

		errors.Assert(o.Argument == nil, "Expected nil ArgumentDescriptor in Specification.")

		o.AddDepth(testDepthDefault)

		errors.Assert(o.Argument != nil, "Expected nil ArgumentDescriptor in Specification.")

		errors.Assert(o.Argument[depthArgLong].FlagId == FlagDepth,
			fmt.Sprintf("(%s) expected (%d)  FlagId:%d", depthArgLong, FlagDepth,
				o.Argument[depthArgLong].FlagId))

		errors.Assert(o.Argument[depthArgLong].Type == Integer,
			fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[depthArgLong].Type))

		errors.Assert(o.Argument[depthArgLong].Default == strconv.Itoa(testDepthDefault),
			fmt.Sprintf("Default should be empty for %s", depthArgLong))

		errors.Assert(o.Argument[depthArgLong].Help == depthHelpText, "help text mismatch")

		errors.Assert(o.Argument[depthArgLong].Parse != nil, "Expect non-nil function pointer")

		errors.Assert(o.Argument[depthArgLong].Expects == ExpectValue, "Next expected should be ExpectValue")

		val, err := strconv.Atoi(o.Argument[depthArgLong].Default)
		if err != nil {
			panic(err)
		}
		if val != testDepthDefault {
			panic(fmt.Sprintf("Default should be false for depth flag (%d): '%d'", FlagDepth, val))
		}
	}()
}
