package facility_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"testing"
)

func TestLogFacility(t *testing.T){
	var f LogFacility.Facility = "test"
	errors.Assert(f == "test", "expected matching string.")
}
