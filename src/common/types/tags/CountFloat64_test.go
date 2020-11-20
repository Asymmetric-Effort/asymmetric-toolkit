package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_CountFloat64(t *testing.T) {
	tag := NewFloat64()
	errors.Assert(tag.Count() == 0, fmt.Sprintf("expect %d (init)", 0))

	for i := 1; i < 10; i++ {
		tag.Add(fmt.Sprintf("test%d", i), float64(i))
		errors.Assert(tag.Count() == i, fmt.Sprintf("expect %d", i))
	}
}
