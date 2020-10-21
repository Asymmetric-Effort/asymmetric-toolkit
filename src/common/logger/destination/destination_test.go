package destination

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDestination(t *testing.T) {
	var ld LogDestination
	errors.Assert(ld == 0, "Expect 0")
	errors.Assert(ld == Stdout, "Expect stdout")

	errors.Assert(Stdout == 0, "Expect 0")
	errors.Assert(File == 1, "Expect 1")
	errors.Assert(Syslog == 2, "Expect 2")
}
