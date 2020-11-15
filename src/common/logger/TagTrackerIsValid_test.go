package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTagTracker_IsValid(t *testing.T) {
	var o TagTracker
	var tagId TagId
	var nextTagId TagId

	func() { // Setup
		tagId = o.Create("test0")
		errors.Assert(tagId == tagId, "expected TagId 0")
		errors.Assert(int(o.tagNames["test0"]) == 0, "expect tagNames")
		errors.Assert(o.tagIds[tagId] == "test0", "expect tagIds")
		errors.Assert(o.IsValid(tagId), "Expected IsValid(0) is true")
	}()

	func() {
		nextTagId = o.Create("test1")
		errors.Assert(nextTagId == nextTagId, "expected TagId 1")
		errors.Assert(o.IsValid(nextTagId), "Expected IsValid(1) is true")
	}()

	func() { // Test
		o.Close(tagId) //Close tagId 0
		errors.Assert(!o.IsValid(tagId), "Expected IsValid(0) is false")
		errors.Assert(o.IsValid(nextTagId), "Expected IsValid(1) is true")
		o.Close(nextTagId) //Close tagId 1
		errors.Assert(!o.IsValid(nextTagId), "Expected IsValid(1) is true")
	}()
}
