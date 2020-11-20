package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestBoolean_Find(t *testing.T){
	var tag Boolean = NewBoolean()
	tag.Add("test1",true)
	errors.Assert(tag["test1"],"expect true")
	tag.Add("test2",false)
	errors.Assert(!tag["test2"],"expect false")
	errors.Assert(tag.Find("test1"), "Expect true")
	errors.Assert(tag.Find("test2"), "Expect true")
	errors.Assert(tag.Delete("test1"), "Expect delete true")
	errors.Assert(!tag.Find("test1"), "Expect false")
	errors.Assert(tag.Find("test2"), "Expect true")

}
