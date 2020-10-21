package destination_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"testing"
)

func TestDestination(t *testing.T) {
	var ld destination.LogDestination
	errors.Assert(ld == 0, "Expect 0")
	errors.Assert(ld == destination.Stdout, "Expect stdout")

	errors.Assert(destination.Stdout == 0, "Expect 0")
	errors.Assert(destination.File == 1, "Expect 1")
	errors.Assert(destination.Syslog == 2, "Expect 2")
}
