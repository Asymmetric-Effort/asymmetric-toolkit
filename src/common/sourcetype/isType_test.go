package sourcetype_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"fmt"
	"testing"
)

func TestDataSource_IsType(t *testing.T) {
	var d sourcetype.DataSource
	errors.Assert(d.IsNotSet(), "Expected NotSet")

	var tests = []struct {
		input      string
		isDict     bool
		isRandom   bool
		isSequence bool
	}{
		{
			"Dictionary",
			true,
			false,
			false,
		}, {
			"Sequence",
			false,
			false,
			true,
		}, {
			"Random",
			false,
			true,
			false,
		},
	}
	for i, test := range tests {
		d.Set(test.input)
		errors.Assert(d.IsDictionary() && test.isDict, fmt.Sprintf("Test %d: Expected Dictionary", i))
		errors.Assert(d.IsSequence() && test.isSequence, fmt.Sprintf("Test %d:Expected Sequence", i))
		errors.Assert(d.IsRandom() && test.isRandom, fmt.Sprintf("Test %d:Expected Random", i))
	}
	defer func() { recover() }()
	d.Set("BadInput")
	t.Error("Expected failure on bad input.")
}

