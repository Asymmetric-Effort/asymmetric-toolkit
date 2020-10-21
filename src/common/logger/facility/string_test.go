package facility_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"testing"
)

func TestLogFacilityString_Happy(t *testing.T) {
	var f LogFacility.Facility = LogFacility.Facility("MyLog") //Happy
	errors.Assert(string(f) == "MyLog", "Expected value not properly stored.")
}

func TestLogFacilityString_Bad(t *testing.T) {
	var f LogFacility.Facility = LogFacility.Facility("MyBadLog!")
	defer func() { recover() }()
	_ = f.Get()
	t.Errorf("Expected error")
}

func TestLogFacilityString_Empty(t *testing.T) {
	var f LogFacility.Facility = LogFacility.Facility("MyBadLog!")
	defer func() { recover() }()
	_ = f.Get()
	t.Errorf("Expected error")
}
