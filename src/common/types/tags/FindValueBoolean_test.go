package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestBoolean_FindValue(t *testing.T){
	var tag Boolean = NewBoolean()
	tag.Add("test1",true)
	errors.Assert(tag["test1"],"expect true")
	tag.Add("test2",false)

	errors.Assert(!tag["test2"],"expect false")
	errors.Assert(!tag["test2"],"expect false")
	errors.Assert(tag.Find("test1"), "Expect true")
	errors.Assert(tag.Find("test2"), "Expect true")

	fmt.Println("tag:",tag)

	errors.Assert(tag.FindValue("test1",true), "findValue() expect true on true")
	errors.Assert(!tag.FindValue("test1",false), "findValue() expect false on false")
	errors.Assert(!tag.FindValue("test2",true), "findValue() expect false on true")
	errors.Assert(tag.FindValue("test2",false), "findValue() expect true on false")

	errors.Assert(tag.Delete("test1"), "Expect delete true")
	errors.Assert(!tag.FindValue("test1",true), "findValue() expect true on true")
	errors.Assert(tag.FindValue("test2",false), "findValue() expect true on false")
	errors.Assert(tag.Delete("test2"), "Expect delete true")
	errors.Assert(!tag.FindValue("test2",false), "findValue() expect true on false")
}
