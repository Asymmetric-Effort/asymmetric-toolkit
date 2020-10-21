package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var o types.List
	o.Set("a,b,c")
	errors.Assert(o.Count() == 3, fmt.Sprintf("Count mismatch.  Expect 3, Encountered %d", o.Count()))
	errors.Assert(o.Get(0)=="a", fmt.Sprintf("Get returned bad value (%s) expected 'a'",o.Get(0)))
	errors.Assert(o.Get(1)=="b", fmt.Sprintf("Get returned bad value (%s) expected 'b'",o.Get(0)))
	errors.Assert(o.Get(2)=="c", fmt.Sprintf("Get returned bad value (%s) expected 'c'",o.Get(0)))
}


func TestListEmpty(t *testing.T) {
	var o types.List
	o.Set("")
	errors.Assert(o.Count() == 1, fmt.Sprintf("Expect 1 (empty), Encountered %d (%s)", o.Count(),o.Get(0)))
}


func TestListPanicOutOfBounds(t *testing.T){
	var o types.List
	o.Set("a,b,c")
	defer func(){recover()}()
	_ = o.Get(3)
	t.Fatal("Expected panic because index is out of bounds.")
}

func TestListString(t *testing.T){
	var o types.List
	var ListString = "a,b,c"
	o.Set(ListString)
	errors.Assert(o.String()==ListString, "Mismatch in List::string()")
}