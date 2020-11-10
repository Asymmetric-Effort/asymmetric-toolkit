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

	if helpArgShort != "h" {
		panic("helpArgShort mismatch")
	}

	if o.Argument != nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	o.AddHelp()

	if o.Argument == nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	for i, argFlag := range []string{helpArgShort, helpArgLong} {

		if o.Argument[argFlag].FlagId != FlagHelp {
			panic(fmt.Sprintf("%d(%s) expected.  FlagId:%d", i, argFlag, o.Argument[argFlag].FlagId))
		}

		if o.Argument[argFlag].Type != None {
			panic(fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[argFlag].Type))
		}

		if o.Argument[argFlag].Default != usageDefault {
			panic(fmt.Sprintf("Default should be empty for %s", argFlag))
		}

		if o.Argument[argFlag].Help != usageHelpText {
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
