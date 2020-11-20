package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFloat64_Add(t *testing.T){
	var tag Float64 = NewFloat64()
	tag.Add("test1",1.0)
	errors.Assert(tag["test1"]==1.0,"expect true")

	tag.Add("test2",2.0)
	errors.Assert(tag["test2"]==2.0,"expect false")
}
