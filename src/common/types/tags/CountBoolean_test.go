package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_CountBoolean(t *testing.T) {
	tag := NewBoolean()
	errors.Assert(tag.Count() == 0, fmt.Sprintf("expect %d (init)", 0))

	for i := 1; i < 10; i++ {
		tag.Add(fmt.Sprintf("test%d", i),true)
		errors.Assert(tag.Count() == i, fmt.Sprintf("expect %d", i))
	}
}
