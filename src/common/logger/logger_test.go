package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogger(t *testing.T){
	var o Logger
	errors.Assert(o.Level == Critical, "Expect critical(0)")
	errors.Assert(o.Destination == Stdout, "expect stdout(0)")
	errors.Assert(o.Settings == nil, "expect nil")
	errors.Assert(o.Writer == nil, "expect nil")
	errors.Assert(&o.tags != nil, "expect non-nil address")
}