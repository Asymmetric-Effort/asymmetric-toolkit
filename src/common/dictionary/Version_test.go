package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionary_Version(t *testing.T) {
	var o Version
	errors.Assert(o == 0, "expect 0")
	errors.Assert(FormatVersion == 0, "Expect 0")
	errors.Assert(ScoreVersion == 0, "Expect 0")
}
