package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTagTracker_Close(t *testing.T) {
	var o TagTracker
	var tagId TagId
	var nextTagId TagId

	func() { // Setup
		tagId = o.Create("test0")
		errors.Assert(tagId == 0, "expected TagId 0")
		errors.Assert(int(o.tagNames["test0"]) == 0, "expect tagNames")
		errors.Assert(o.tagIds[0] == "test0", "expect tagIds")
	}()

	func() {
		nextTagId = o.Create("test1")
		errors.Assert(nextTagId == 1, "expected TagId 1")
	}()

	func() { // Test
		o.Close(tagId) //Close tagId 0
		defer func() { recover() }()
		errors.Assert(tagId == 0, "expected TagId 0")
		errors.Assert(int(o.tagNames["test0"]) == 0, "expect error because this should be deleted.")
		errors.Assert(o.tagIds[0] == "test0", "expect error because this should be deleted.")
	}()

	func() { //Test second delete
		defer func() { recover() }()
		o.Close(tagId)
		panic("Test second delete expected error.")
	}()

	func() { // Make 2nd tag global
		o.AddGlobal(nextTagId)
		errors.Assert(o.global[nextTagId] == struct{}{}, "Expect 2nd tag to be global.")
	}()

	func() { // Test
		o.Close(nextTagId) //Close tagId 1
		defer func() { recover() }()
		errors.Assert(nextTagId == 1, "expected TagId 1")
		errors.Assert(o.tagNames["test1"] == nextTagId, "expect error because this should be deleted.")
		errors.Assert(o.tagIds[nextTagId] == "test1", "expect error because this should be deleted.")
		errors.Assert(o.global[nextTagId] == struct{}{}, "expect error.")
	}()
}
