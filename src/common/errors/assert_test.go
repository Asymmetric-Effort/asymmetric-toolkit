package errors_test

import "testing"

func TestAssert(t *testing.T){
	defer func(){recover()}()
	Assert(false,"expect true") //cause assertion error."
	t.FailNow()
}