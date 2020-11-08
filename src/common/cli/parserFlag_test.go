package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestParserFlag(t *testing.T) {
	func() { //Test flag present, value should be true
		parser := ParserFlag("flag")
		inp := "flag"
		err, val := parser(&inp)
		if err != nil {
			panic(err)
		}
		errors.Assert(val.Type == Boolean, "Expected boolean result")
		errors.Assert(val.String() == "true", "expected true (string)")
		errors.Assert(val.Boolean(), "expected true")
	}()
	func() { //Test flag not present, value should be false
		parser := ParserFlag("flag")
		inp := ""
		err, val := parser(&inp)
		if err != nil {
			panic(err)
		}
		errors.Assert(val.Type == Boolean, "Expected boolean result")
		errors.Assert(val.String() == "false", "expected false (string)")
		errors.Assert(!val.Boolean(), "expected false")
	}()

}
