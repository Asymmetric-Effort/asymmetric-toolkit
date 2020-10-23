package LogFacility

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestLogFacility(t *testing.T){
	var f Facility = "test"
	errors.Assert(f == "test", "expected matching string.")
}
