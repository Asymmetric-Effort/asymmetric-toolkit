package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFloat64_Find(t *testing.T) {
	var tag Float64 = NewFloat64()

	tag.Add("test1", 1.1)
	errors.Assert(tag["test1"] == 1.1, "expect true")

	tag.Add("test2", 2.2)
	errors.Assert(tag["test2"] == 2.2, "expect false")

	errors.Assert(tag.Find("test1"), "Expect true")
	errors.Assert(tag.Find("test2"), "Expect true")
	errors.Assert(tag.Delete("test1"), "Expect delete true")
	errors.Assert(!tag.Find("test1"), "Expect false")
	errors.Assert(tag.Find("test2"), "Expect true")

}
