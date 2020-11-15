package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogger_TagTable(t *testing.T){
	var o TagTable = make(TagTable,1)
	o[0]=struct{}{}
	errors.Assert(o[0]==struct{}{}, "expect non-nil")
}