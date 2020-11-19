package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionary_Configuration (t *testing.T){
	var o Configuration
	errors.Assert(!o.Overwrite, "expected false")
	errors.Assert(o.FormatVersion == 0, "expected 0")
	errors.Assert(o.FileName == "", "expect empty string.")
	errors.Assert(o.ScoreVersion == 0, "expect 0")
}