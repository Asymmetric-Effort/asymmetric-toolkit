package counter_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/counter"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCounter(t *testing.T) {
	var c counter.Counter
	errors.Assert(c.String() == "", "Expect an empty string if we are not initialized.")
}
