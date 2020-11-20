package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestBoolean_Get(t *testing.T) {
	var tag Boolean = NewBoolean()
	tag.Add("test1", true)
	errors.Assert(tag.Get("test1"), "expect true")
	tag.Add("test2", false)
	errors.Assert(!tag.Get("test2"), "expect false")
}
