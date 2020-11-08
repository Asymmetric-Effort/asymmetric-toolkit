package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"strings"
	"testing"
)

func TestParserBool(t *testing.T) {
	parser := ParserBool()
	func() {
		inp := "true"
		err, val := parser(&inp)
		if err != nil {
			t.Errorf("Expected Boolean(%v):%v", inp, val)
		}
		errors.Assert(val.Type == Boolean,"Expected boolean type")
		errors.Assert(val.String()=="true","Expected boolean true string")
		errors.Assert(val.Boolean(), "Expected true")
		defer func(){recover()}()
		errors.Assert(val.Float()==0.0, "Expected error")
		errors.Assert(val.Float()!=0.0, "Expected error")
		errors.Assert(val.Integer()==0.0, "Expected error")
	}()
	func() {
		testData := []string{"true", "false", "True", "False", "TrUe", "faLse", "true ", " false"}
		for _, inp := range testData {
			err, val := parser(&inp)
			if err != nil {
				panic(err)
			}
			if val.Type != Boolean {
				t.Error("Expected Boolean(1)")
			}
			if val.Value != strings.TrimSpace(strings.ToLower(inp)) {
				t.Errorf("Expected matching value (1)")
			}
		}
	}()

	func() {
		defer func() { recover() }()
		testData := []string{"1 ", " 2", "BadInput", "", "<", "'", ">>", "5534523$$$"}
		for _, inp := range testData {
			err, val := parser(&inp)
			if err != nil {
				panic(err)
			}
			if val.Type != Boolean {
				t.Error("Expected Boolean(2)")
			}
			if val.Value != strings.TrimSpace(strings.ToLower(inp)) {
				t.Errorf("Expected matching value (2)")
			}
		}
	}()
}
