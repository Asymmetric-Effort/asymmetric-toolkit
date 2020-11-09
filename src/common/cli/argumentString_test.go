package cli

import "testing"

func TestArgument_String(t *testing.T) {
	func() {
		const ExpectedValue = "Foo"
		var o Argument
		o.Type = String
		o.Value = ExpectedValue
		if o.Value != ExpectedValue {
			t.Error("Expected String(1).")
		}
		if o.String() != ExpectedValue {
			t.Error("Expected String(2)")
		}
	}()
	func() {
		const ExpectedValue = "A,B,C,D"
		var o Argument
		o.Type = List
		o.Value = "1,"+ExpectedValue
		if o.String() != ExpectedValue {
			t.Errorf("Expected String List: %s",o.String())
		}
	}()
}
