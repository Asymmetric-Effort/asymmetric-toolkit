package utils

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCountNumberFrequency(t *testing.T) {
	n := 12345
	c := *CountNumberFrequency(n)
	errors.Assert(c[0] == 0, "Expect 0 0s")
	errors.Assert(c[1] == 1, "Expect 1 1s")
	errors.Assert(c[2] == 1, "Expect 1 2s")
	errors.Assert(c[3] == 1, "Expect 1 3s")
	errors.Assert(c[4] == 1, "Expect 1 4s")
	errors.Assert(c[5] == 1, "Expect 1 5s")

	n = 11111
	c = *CountNumberFrequency(n)
	errors.Assert(c[0] == 0, "Expect 0 0s")
	errors.Assert(c[1] == 5, "Expect 5 1s")
	errors.Assert(c[2] == 0, "Expect 0 2s")
	errors.Assert(c[3] == 0, "Expect 0 3s")
	errors.Assert(c[4] == 0, "Expect 0 4s")
	errors.Assert(c[5] == 0, "Expect 0 5s")

	n = 122333444455555
	c = *CountNumberFrequency(n)
	errors.Assert(c[0] == 0, "Expect 0 0s")
	errors.Assert(c[1] == 1, "Expect 1 1s")
	errors.Assert(c[2] == 2, "Expect 2 2s")
	errors.Assert(c[3] == 3, "Expect 3 3s")
	errors.Assert(c[4] == 4, "Expect 4 4s")
	errors.Assert(c[5] == 5, "Expect 5 5s")

}
