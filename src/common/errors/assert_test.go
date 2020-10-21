package errors_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestAssert(t *testing.T){
	defer func(){recover()}()
	errors.Assert(false,"expect true") //cause assertion error."
	t.FailNow()
}