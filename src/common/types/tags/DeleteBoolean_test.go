package tags

import (
"asymmetric-effort/asymmetric-toolkit/src/common/errors"
"testing"
)

func TestBoolean_Delete(t *testing.T){
	var tag Boolean = NewBoolean()
	tag.Add("test1",true)
	errors.Assert(tag["test1"],"expect true")
	errors.Assert(tag.Delete("test1"), "expect true")
	errors.Assert(!tag.Delete("test1"), "expect false")
	errors.Assert(!tag.Find("test1"), "expect false find after delete")
}
