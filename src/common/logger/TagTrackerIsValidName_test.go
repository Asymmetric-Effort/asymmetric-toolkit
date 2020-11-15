package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTagTracker_isValidName(t *testing.T){
	var o TagTracker
	var tagId TagId
	var nextTagId TagId
	var tagName string= "test0"
	var nextTagName string= "test1"

	func() { // Setup
		tagId = o.Create(tagName)
		errors.Assert(tagId == tagId, "expected TagId 0")
		errors.Assert(int(o.tagNames["test0"]) == 0, "expect tagNames")
		errors.Assert(o.tagIds[tagId] == "test0", "expect tagIds")
		errors.Assert(o.IsValidName(tagName),"expect tagName test0")
	}()

	func() {
		nextTagId = o.Create(nextTagName)
		errors.Assert(o.IsValidName(nextTagName),"expect tagName test1")
	}()

	func() { // Test
		o.Close(tagId) //Close tagId 0
		errors.Assert(!o.IsValidName(tagName),"expect no tagName test0")
		errors.Assert(o.IsValidName(nextTagName),"expect tagName test1")
		o.Close(nextTagId) //Close tagId 1
		errors.Assert(!o.IsValidName(nextTagName),"expect no tagName test1")
	}()
}