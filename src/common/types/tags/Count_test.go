package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_Count(t *testing.T) {
	tag := NewTag()
	errors.Assert(tag.Count() == 0, fmt.Sprintf("expect %d (init)", 0))

	for i := 1; i < 10; i++ {
		tag.Add(fmt.Sprintf("test%d", i))
		errors.Assert(tag.Count() == i, fmt.Sprintf("expect %d", i))
	}
}
