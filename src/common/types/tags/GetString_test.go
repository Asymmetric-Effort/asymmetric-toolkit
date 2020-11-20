package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestString_Get(t *testing.T) {
	var tag String = NewString()
	tag.Add("test1", "MyString")
	errors.Assert(tag.Get("test1") == "MyString", "expect true")
}
