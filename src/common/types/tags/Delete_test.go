package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTag_Delete(t *testing.T) {
	o := NewTag()
	o.Add("test")
	errors.Assert(o["test"]==struct{}{}, "[1] value mismatch")
	o.Delete("test")
	errors.Assert(!o.Find("test"), "[1] tag not found")
}
