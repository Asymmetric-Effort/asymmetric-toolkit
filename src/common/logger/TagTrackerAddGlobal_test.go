package logger

import "testing"

func TestLogger_TagTracker_AddGlobal(t *testing.T) {
	func() { // Expect failure creating a global tag that hasn't been created.
		var o TagTracker
		defer func() { recover() }()
		o.AddGlobal(0)
		panic("Expected failure on AddGlobal() for tagId not created.")
	}()
	func(){
		var o TagTracker
		tagId:=o.Create("test")
		o.AddGlobal(tagId)
	}()
}
