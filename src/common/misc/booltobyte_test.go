package misc

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestBoolToByte(t *testing.T){
	errors.Assert(BoolToByte(true)==1, "expect 1")
	errors.Assert(BoolToByte(false)==0, "expect 1")
}