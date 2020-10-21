package counter_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/counter"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCounter_String(t *testing.T) {
	var c counter.Counter
	const charset = "01"
	const wordSize = 2
	c.Setup(charset, wordSize)
	errors.Assert(c.String() == "00", "Expect 00 in initial state.")
	c.Increment(0)
	errors.Assert(c.String() == "01", "Expect 00 in initial state.")
	c.Increment(0)
	errors.Assert(c.String() == "10", "Expect 00 in initial state.")
	c.Increment(0)
	errors.Assert(c.String() == "11", "Expect 00 in initial state.")
}
