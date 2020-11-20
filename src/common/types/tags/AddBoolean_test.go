package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestBoolean_Add(t *testing.T){
	o := NewBoolean()
	for i := 1; i < 255; i++ {
		tagName := fmt.Sprintf("test%d", i)
		o.Add(tagName, true)
		errors.Assert(o[tagName],"expect true")
	}
	defer func(){recover()}()
	o.Add("Too many",true)
	o.Add("Too many",true)
	panic("Expected error for too many.  It didn't happen.")
}
