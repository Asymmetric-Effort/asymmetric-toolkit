package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestString_Add(t *testing.T){
	var tag String = NewString()

	tag.Add("test1","myString1")

	errors.Assert(tag["test1"]=="myString1","expect true")

	tag.Add("test2","myString2")

	errors.Assert(tag["test2"]=="myString2","expect false")
}
