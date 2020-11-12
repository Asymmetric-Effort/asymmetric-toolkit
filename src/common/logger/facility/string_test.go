package logFacility

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogFacilityString_Happy(t *testing.T){
	var f Facility
	f= Facility("MyLog") //Happy
	errors.Assert(string(f)=="MyLog", "Expected value not properly stored.")
}

func TestLogFacilityString_Bad(t *testing.T){
	var f Facility
	defer func(){recover()}()
	f= Facility("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}

func TestLogFacilityString_Empty(t *testing.T){
	var f Facility
	defer func(){recover()}()
	f= Facility("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}
