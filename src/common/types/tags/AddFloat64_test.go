package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestFloat64_Add(t *testing.T){
	o:= NewFloat64()
	for i := 1; i < 255; i++ {
		tagName := fmt.Sprintf("test%d", i)
		o.Add(tagName, 1.0)
		errors.Assert(o[tagName] == 1.0, "expect true")
	}
	defer func(){recover()}()
	o.Add("Too many",1.0)
	o.Add("Too many",2.0)
	panic("Expected error for too many.  It didn't happen.")
}
