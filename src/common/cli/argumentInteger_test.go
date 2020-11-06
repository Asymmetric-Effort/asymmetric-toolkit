package cli

import "testing"

func TestArgument_Integer(t *testing.T) {
	func() { // Happy Path tests
		var o Argument
		o.Type = Integer

		o.Value = "0"
		if o.Float() != 0 {
			t.Error("Expected parsed result")
		}
		o.Value = "1"
		if o.Float() != 1 {
			t.Error("Expected parsed result")
		}
		o.Value = "-1"
		if o.Float() != -1 {
			t.Error("Expected parsed result")
		}
	}()

	func() { //Sad-Path tests
		var o Argument
		defer func() { recover() }()
		o.Type = Integer
		o.Value = "1"
		_ = o.Float()
		o.Value = "1.0"
		_ = o.Float()
		o.Value = "foo"
		_ = o.Float()
		o.Value = ""
		_ = o.Float()
		o.Value = "true"
		_ = o.Float()
	}()

}
