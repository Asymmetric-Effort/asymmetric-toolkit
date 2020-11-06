package cli

import "testing"

func TestArgument_String(t *testing.T) {
	const ExpectedValue = "Foo"
	var o Argument
	o.Type = String
	o.Value=ExpectedValue
	if o.Value != ExpectedValue {
		t.Error("Expected String.")
	}
}