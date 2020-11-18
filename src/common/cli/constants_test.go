package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCliConstants(t *testing.T) {
	errors.Assert(ErrSuccess == 0, "Expected 0")
	errors.Assert(ErrArgumentParseError == 1, "Expected 1")

	errors.Assert(ErrBadFilename == 10, "Expected 10")
	errors.Assert(ErrFileNotFound == 11, "Expected 11")
	errors.Assert(ErrFileOpenFailed == 12, "Expected 12")
	errors.Assert(ErrFileRead == 13, "Expected 13")

	errors.Assert(MissingArgumentsMessage == "Missing arguments (use --help for usage)",
		"Expected string mismatch")
}
