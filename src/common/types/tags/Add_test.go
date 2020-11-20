package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_Add(t *testing.T) {
	o := NewTag()

	for i := 1; i < 255; i++ {
		tagName := fmt.Sprintf("test%d", i)
		o.Add(tagName)
		errors.Assert(o[tagName] == struct{}{}, fmt.Sprintf("[%s] value mismatch", tagName))
	}
	defer func(){recover()}()
	o.Add("Too many")
	o.Add("Too many")
	panic("Expected error for too many.  It didn't happen.")
}
