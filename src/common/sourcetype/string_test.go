package sourcetype_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"testing"
)

func TestDataSource_String(t *testing.T) {
	var o sourcetype.DataSource
	o = sourcetype.Random
	errors.Assert(o.String() == "Random", "Expected Random")
	o = sourcetype.Sequence
	errors.Assert(o.String() == "Sequence", "Expected Sequence")
	o = sourcetype.Dictionary
	errors.Assert(o.String() == "Dictionary", "Expected Dictionary")
}

func TestDataSource_String_Bad(t *testing.T) {
	var o sourcetype.DataSource = BadDataSource
	defer func(){recover()}()
	_ = o.String()
	t.FailNow()
}
