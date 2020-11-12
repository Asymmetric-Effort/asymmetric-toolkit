package logFacility

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogFacility(t *testing.T){
	var f Facility = "test"
	errors.Assert(f == "test", "expected matching string.")
}
