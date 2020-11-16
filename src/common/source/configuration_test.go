package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestSource_Configuration(t *testing.T) {
	var config Configuration
	errors.Assert(config.Dictionary == "", "expected empty string")
	errors.Assert(config.MaxWordCount == 0, "expected 0")
	errors.Assert(config.MinWordSize == 0, "expected 0")
	errors.Assert(config.MaxWordSize == 0, "expected 0")
	errors.Assert(config.Pattern == "", "expected empty string")
	errors.Assert(config.Mode == "", "expected empty string")
}