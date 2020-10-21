package facility_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"testing"
)

func TestLogFacilityGet_Happy(t *testing.T) {
	const facilityName = "MyLog"
	var f LogFacility.Facility = LogFacility.Facility(facilityName) //Happy
	errors.Assert(string(f) == facilityName, "Expected value not properly stored.")
}

func TestLogFacilityGet_Bad(t *testing.T) {
	const facilityName = "MyBadLog!"
	var f LogFacility.Facility = LogFacility.Facility(facilityName)
	defer func() { recover() }()
	_ = f.Get()
	t.Errorf("Expected error")
}

func TestLogFacilityGet_Empty(t *testing.T) {
	var f LogFacility.Facility = LogFacility.Facility("")
	defer func() { recover() }()
	_ = f.Get()
	t.Errorf("Expected error")
}
