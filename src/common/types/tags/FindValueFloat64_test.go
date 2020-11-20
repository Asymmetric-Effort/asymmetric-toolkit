package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFloat64_FindValue(t *testing.T){
	var tag Float64 = NewFloat64()
	tag.Add("test1",1.1)
	tag.Add("test2",2.2)

	errors.Assert(tag["test1"]==1.1,"expect true")
	errors.Assert(tag["test2"]==2.2,"expect true")
	errors.Assert(tag.Find("test1"), "Expect true")
	errors.Assert(tag.Find("test2"), "Expect true")

	errors.Assert(tag.FindValue("test1",1.1), "findValue() expect true on true")
	errors.Assert(!tag.FindValue("test1",2.2), "findValue() expect false on false")
	errors.Assert(!tag.FindValue("test2",3.3), "findValue() expect false on true")
	errors.Assert(tag.FindValue("test2",2.2), "findValue() expect true on false")

	errors.Assert(tag.Delete("test1"), "Expect delete true")
	errors.Assert(!tag.FindValue("test1",1.1), "findValue() expect true on true")
	errors.Assert(tag.FindValue("test2",2.2), "findValue() expect true on false")
	errors.Assert(tag.Delete("test2"), "Expect delete true")
	errors.Assert(!tag.FindValue("test2",2.2), "findValue() expect true on false")
}
