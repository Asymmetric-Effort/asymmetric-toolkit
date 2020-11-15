package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTagTracker_DeleteGlobal(t *testing.T) {
	var o TagTracker
	var tagId TagId
	var nextTagId TagId

	func() { // Setup two tags.
		tagId = o.Create("test0")
		errors.Assert(tagId == 0, "expected TagId 0")
		errors.Assert(int(o.tagNames["test0"]) == 0, "expect tagNames")
		errors.Assert(o.tagIds[0] == "test0", "expect tagIds")
		nextTagId = o.Create("test1")
		errors.Assert(nextTagId == 1, "expected TagId 1")
		o.AddGlobal(nextTagId)
		errors.Assert(o.global[nextTagId] == struct{}{}, "Expect 2nd tag to be global.")
	}()

	func() { // Test DeleteGlobal()
		errors.Assert(tagId == 0, "expected TagId 0")
		errors.Assert(nextTagId == 1, "expected nextTagId 1")
		o.DeleteGlobal(tagId) //Delete the global tag tagId 0
		defer func() { recover() }()
		errors.Assert(int(o.tagNames["test0"]) == 0, "expect error because this should be deleted.")
		errors.Assert(o.tagIds[0] == "test0", "expect error because this should be deleted.")

		errors.Assert(o.tagNames["test1"] == nextTagId, "expect error because this should be deleted.")
		errors.Assert(o.tagIds[nextTagId] == "test1", "expect error because this should be deleted.")
		errors.Assert(o.global[nextTagId] == struct{}{}, "expect error.")
	}()
}