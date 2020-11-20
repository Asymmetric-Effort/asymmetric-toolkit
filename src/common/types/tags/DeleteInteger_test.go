package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func BenchmarkInteger_Delete(b *testing.B) {
	var tag Integer = NewInteger()
	tag.Add("test1",1)
	errors.Assert(tag["test1"]==1,"expect true")
	errors.Assert(tag.Delete("test1"), "expect true")
	errors.Assert(!tag.Delete("test1"), "expect false")
	errors.Assert(!tag.Find("test1"), "expect false find after delete")
}