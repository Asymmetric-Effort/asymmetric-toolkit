package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_ShowVersion(t *testing.T) {
	var o Specification

	fmt.Println("Starting TestSpecification_ShowVersion()")

	o.Version = "1.1.1"
	o.ProgramName = "show_version_test"

	inp := "--version"

	output := errors.CaptureStdOut(func() {
		o.ShowVersion(&inp)
	})
	fmt.Printf("--DEBUG:\n"+
		"\tVersion:     '%s'\n"+
		"\tProgramName: '%s'\n"+
		"\tOutput:      '%s'\n"+
		"---DEBUG\n", o.Version, o.ProgramName, output)

	errors.Assert(output == fmt.Sprintf("%s (%s)\n", o.ProgramName, o.Version), "version mismatch")
}
