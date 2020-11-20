package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestInteger_Get(t *testing.T) {
	var tag Integer = NewInteger()
	tag.Add("test1", 1)
	errors.Assert(tag.Get("test1") == 1, "expect true")
}
