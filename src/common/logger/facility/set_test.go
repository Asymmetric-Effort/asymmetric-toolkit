package LogFacility_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogFacilitySet_Happy(t *testing.T){
	var f Facility
	f.Set("MyLog") //Happy
	errors.Assert(string(f)=="MyLog", "Expected value not properly stored.")
}

func TestLogFacilitySet_Bad(t *testing.T){
	var f Facility
	defer func(){recover()}()
	f.Set("MyBadLog!")
	_ = f.Get()
	t.Errorf("Expected error")
}

func TestLogFacilitySet_Empty(t *testing.T){
	var f Facility
	defer func(){recover()}()
	f.Set("")
	_ = f.Get()
	t.Errorf("Expected error")
}
