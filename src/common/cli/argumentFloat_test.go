package cli

import "testing"

func TestArgument_Float(t *testing.T) {
	func() { // Happy Path tests
		var o Argument
		o.Type = Float

		o.Value = "1"
		if o.Float() != 1.0 {
			t.Errorf("Expected parsed result (%s)", o.Value)
		}
		o.Value = "-1"
		if o.Float() != -1.0 {
			t.Errorf("Expected parsed result (%s)", o.Value)
		}
		o.Value = "1.0"
		if o.Float() != 1.0 {
			t.Errorf("Expected parsed result (%s)", o.Value)
		}
		o.Value = "-1.0"
		if o.Float() != -1.0 {
			t.Errorf("Expected parsed result (%s)", o.Value)
		}
		o.Value = "0.1"
		if o.Float() != 0.1 {
			t.Errorf("Expected parsed result (%s)", o.Value)
		}
		o.Value = "0"
		if o.Float() != 0.0 {
			t.Errorf("Expected parsed result (%s)", o.Value)
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
