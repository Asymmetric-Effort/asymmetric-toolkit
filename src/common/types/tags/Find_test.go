package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTag_Find(t *testing.T) {
	o:=NewTag()
	o.Add("test")
	errors.Assert(o.Find("test"),"expect true")
	errors.Assert(!o.Find("BadTest"),"expect false")
}
