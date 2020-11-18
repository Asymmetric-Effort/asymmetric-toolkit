package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestProgramDescription(t *testing.T) {
	errors.Assert(ProgramDescription == "This program is used to create dictionaries from "+
		"some third-party source, such as a text file.",
		"Mismatch")
}
