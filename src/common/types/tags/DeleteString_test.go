package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestString_Delete(t *testing.T) {
	var tag String = NewString()
	tag.Add("test1","MyString")
	errors.Assert(tag["test1"]=="MyString","expect true")
	errors.Assert(tag.Delete("test1"), "expect true")
	errors.Assert(!tag.Delete("test1"), "expect false")
	errors.Assert(!tag.Find("test1"), "expect false find after delete")
}
