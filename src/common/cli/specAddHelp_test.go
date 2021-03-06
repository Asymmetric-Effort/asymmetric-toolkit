package cli

import (
	"fmt"
	"testing"
)

func TestSpecification_AddUsage(t *testing.T) {

	var o Specification

	if helpArgLong != "help" {
		panic("helpArgLong mismatch")
	}

	if o.Argument != nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	o.AddHelp()

	if o.Argument == nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	for i, argFlag := range []string{helpArgLong} {

		if o.Argument[argFlag].FlagId != FlagHelp {
			panic(fmt.Sprintf("%d(%s) expected.  FlagId:%d", i, argFlag, o.Argument[argFlag].FlagId))
		}

		if o.Argument[argFlag].Type != Boolean {
			panic(fmt.Sprintf("Boolean Argument type expected.  Type:%d", o.Argument[argFlag].Type))
		}

		if o.Argument[argFlag].Default != helpDefault {
			panic(fmt.Sprintf("Default should be empty for %s", argFlag))
		}

		if o.Argument[argFlag].Help != helpHelpText {
			panic("usageHelpText mismatch")
		}

		if o.Argument[argFlag].Parse == nil {
			panic("Expect non-nil function pointer")
		}

		if o.Argument[argFlag].Expects != ExpectNone {
			panic("Next expected should be ExpectNone")
		}
	}
}
