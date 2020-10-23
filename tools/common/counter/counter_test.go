package counter

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestCounter (t *testing.T){
	var c Counter
	errors.Assert(c.runes == nil, "expected nil pointer for c.runes")
	errors.Assert(c.data == nil, "expected nil pointer for c.data")
	errors.Assert(c.maxBase == 0, "expected 0 value for c.maxBase")
}