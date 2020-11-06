package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestParserFunction(t *testing.T) {
	var o ParserFunction
	var nullStringPtr *string = nil
	o = func(s *string) (err error, val *Argument) {
		errors.Assert(s == nil, "expect nil input")
		return nil, nil
	}
	a, b := o(nullStringPtr)
	if a == nil && b == nil {
		return
	}
	t.Error("failed to return nil, nil")
}
