package counter_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/counter"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCounter (t *testing.T){
	var c counter.Counter
	errors.Assert(c.Runes == nil, "expected nil pointer for c.runes")
	errors.Assert(c.Data == nil, "expected nil pointer for c.data")
	errors.Assert(c.MaxBase == 0, "expected 0 value for c.maxBase")
}