package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestString_Find(t *testing.T) {
	var tag String = NewString()

	tag.Add("test1", "test")
	errors.Assert(tag["test1"] == "test", "expect true")

	tag.Add("test2", "test2")
	errors.Assert(tag["test2"] == "test2", "expect false")

	errors.Assert(tag.Find("test1"), "Expect true")
	errors.Assert(tag.Find("test2"), "Expect true")
	errors.Assert(tag.Delete("test1"), "Expect delete true")
	errors.Assert(!tag.Find("test1"), "Expect false")
	errors.Assert(tag.Find("test2"), "Expect true")

}
