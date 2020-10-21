package LogFacility_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogFacilityGet_Happy(t *testing.T){
	var f Facility
	f= Facility("MyLog") //Happy
	errors.Assert(string(f)=="MyLog", "Expected value not properly stored.")
}

func TestLogFacilityGet_Bad(t *testing.T){
	var f Facility
	defer func(){recover()}()
	f= Facility("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}

func TestLogFacilityGet_Empty(t *testing.T){
	var f Facility
	defer func(){recover()}()
	f= Facility("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}
