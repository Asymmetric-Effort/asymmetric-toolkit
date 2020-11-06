package cli

import (
	"strings"
	"testing"
)

func TestParserBool(t *testing.T) {
	parser := ParserBool()

	func() {
		testData := []string{"true", "false","True","False","TrUe","faLse","true ", " false"}
		for _, inp := range testData {
			err, val := parser(&inp)
			if err != nil {
				panic(err)
			}
			if val.Type != Boolean{
				t.Error("Expected Boolean(1)")
			}
			if val.Value != strings.TrimSpace(strings.ToLower(inp)) {
				t.Errorf("Expected matching value (1)")
			}
		}
	}()

	func() {
		defer func(){recover()}()
		testData := []string{"1 ", " 2","BadInput","","<","'",">>","5534523$$$"}
		for _, inp := range testData {
			err, val := parser(&inp)
			if err != nil {
				panic(err)
			}
			if val.Type != Boolean{
				t.Error("Expected Boolean(2)")
			}
			if val.Value != strings.TrimSpace(strings.ToLower(inp)) {
				t.Errorf("Expected matching value (2)")
			}
		}
	}()
}
