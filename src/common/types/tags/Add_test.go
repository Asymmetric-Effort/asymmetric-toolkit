package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTag_Add(t *testing.T) {
	o := NewTag()
	o.Add("test")
	errors.Assert(o["test"]==struct{}{}, "[1] value mismatch")
}
