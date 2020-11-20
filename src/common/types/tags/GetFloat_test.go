package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFloat_Get(t *testing.T) {
	var tag Float64 = NewFloat64()
	tag.Add("test1", 1.0)
	errors.Assert(tag.Get("test1") == 1.0, "expect true")
}
