package cli

import (
	"fmt"
	"testing"
)

func TestSpecification_AddVersion(t *testing.T) {

	var o Specification

	if versionHelpText != "Show version string" {
		panic("versionHelpText mismatch")
	}
	if versionDefault != "" {
		panic("versionDefault mismatch")
	}
	if versionArgLong != "version" {
		panic("version flag (long) mismatch")
	}
	if versionArgShort != "v" {
		panic("version flag (short) mismatch")
	}

	if o.Argument != nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	o.AddVersion()

	if o.Argument == nil {
		panic("Expected nil ArgumentDescriptor in Specification.")
	}

	for i, argFlag := range []string{versionArgShort, versionArgLong} {

		if o.Argument[argFlag].FlagId != flagVersion {
			panic(fmt.Sprintf("%d(%s) expected.  FlagId:%d", i, argFlag, o.Argument[argFlag].FlagId))
		}

		if o.Argument[argFlag].Type != None {
			panic(fmt.Sprintf("String Argument type expected.  Type:%d", o.Argument[argFlag].Type))
		}

		if o.Argument[argFlag].Default != versionDefault {
			panic(fmt.Sprintf("Default should be empty for %s", argFlag))
		}

		if o.Argument[argFlag].Help != versionHelpText {
			panic("versionHelpText mismatch")
		}

		if o.Argument[argFlag].Parse == nil {
			panic("Expect non-nil function pointer")
		}

		if o.Argument[argFlag].Expects != ExpectEnd {
			panic("Next expected should be ExpectEnd")
		}
	}
}
