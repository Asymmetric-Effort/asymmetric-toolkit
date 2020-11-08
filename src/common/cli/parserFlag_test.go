package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestParserFlag(t *testing.T) {
	parser := ParserFlag()
	func() {
		inp := "--flag"
		err, val := parser(&inp)
		if err != nil {
			panic(err)
		}
		errors.Assert(val.Type == Boolean, "Expected boolean result")
		errors.Assert(val.String() == "true", "expected true (string)")
		errors.Assert(val.Boolean(), "expected true")
	}()
}
