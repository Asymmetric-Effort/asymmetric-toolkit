package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFloat64_Delete(t *testing.T) {
	var tag Float64 = NewFloat64()
	tag.Add("test1",1.0)
	errors.Assert(tag["test1"]==1.0,"expect true")
	errors.Assert(tag.Delete("test1"), "expect true")
	errors.Assert(!tag.Delete("test1"), "expect false")
	errors.Assert(!tag.Find("test1"), "expect false find after delete")
}