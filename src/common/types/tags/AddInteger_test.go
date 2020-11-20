package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestInteger_Add(t *testing.T){
	var tag Integer = NewInteger()
	tag.Add("test1",1)
	errors.Assert(tag["test1"]==1,"expect true")

	tag.Add("test2",2)
	errors.Assert(tag["test2"]==2,"expect false")
}
