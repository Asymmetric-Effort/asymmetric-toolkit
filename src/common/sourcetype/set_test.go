package sourcetype_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"strconv"
	"testing"
)

func TestDataSource_GetSet(t *testing.T) {
	var o sourcetype.DataSource
	var tests = []struct {
		input    string
		expected sourcetype.DataSource
	}{
		{
			"Dictionary",
			sourcetype.Dictionary,
		}, {
			"Sequence",
			sourcetype.Sequence,
		}, {
			"Random",
			sourcetype.Random,
		},
	}
	errors.Assert(o == sourcetype.NotSet, "Expected NotSet")
	errors.Assert(o.Get() == sourcetype.NotSet, "Expected NotSet")
	for _, test := range tests {
		o.Set(test.input)
		errors.Assert(o == test.expected, "Expected "+test.input)
		errors.Assert(o.Get() == test.expected, "Expected "+test.input)
		errors.Assert(o.String() == test.input, "Expected "+test.input)
	}
	go func() {
		var o1 sourcetype.DataSource
		defer func() { recover() }()
		o1.Set(strconv.Itoa(int(BadDataSource)))
		t.FailNow()
	}()
	go func() {
		//Sad Path
		var o1 sourcetype.DataSource = BadDataSource
		defer func() { recover() }()
		_ = o1.Get()
		t.FailNow()
	}()
	go func() {
		var o1 sourcetype.DataSource = BadDataSource
		defer func() { recover() }()
		_ = o1.String()
		t.FailNow()
	}()
}
