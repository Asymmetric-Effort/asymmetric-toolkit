package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_SerializeBoolean(t *testing.T) {
	o := NewBoolean()
	errors.Assert(o.Count() == 0, "expect 0")
	o.Add("test", true)
	errors.Assert(o.Count() == 1, "expect 1")
	errors.Assert(o.Find("test"), "expect true")
	errors.Assert(o.Delete("test"), "expect true")
	o.Add("test1", true)
	s := o.SerializeBoolean()
	errors.Assert(s != nil, fmt.Sprintf("expect not nil  V:%v",s))
	o.Add("test2", true)
	o.Add("test3", true)
	o.Add("test4", true)
	errors.Assert(o.Count() == 4, "expect 4")
	s = o.SerializeBoolean()
}
