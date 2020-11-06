package cli

import "testing"

func TestArgument_Integer(t *testing.T) {
	func() { // Happy Path tests
		var o Argument
		o.Type = Integer

		o.Value = "0"
		if o.Integer() != 0 {
			t.Error("Expected parsed result")
		}
		o.Value = "1"
		if o.Integer() != 1 {
			t.Error("Expected parsed result")
		}
		o.Value = "-1"
		if o.Integer() != -1 {
			t.Error("Expected parsed result")
		}
	}()

	func() { //Sad-Path tests
		var o Argument
		defer func() { recover() }()
		o.Type = Integer
		o.Value = "1"
		_ = o.Integer()
		o.Value = "1.0"
		_ = o.Integer()
		o.Value = "foo"
		_ = o.Integer()
		o.Value = ""
		_ = o.Integer()
		o.Value = "true"
		_ = o.Integer()
	}()

}
