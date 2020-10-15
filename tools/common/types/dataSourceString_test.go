package types

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestDataSource_String(t *testing.T) {
	var o DataSource
	o = Random
	errors.Assert(o.String() == "Random", "Expected Random")
	o = Sequence
	errors.Assert(o.String() == "Sequence", "Expected Sequence")
	o = Dictionary
	errors.Assert(o.String() == "Dictionary", "Expected Dictionary")
}

func TestDataSource_String_Bad(t *testing.T) {
	var o DataSource = BadDataSource
	defer func(){recover()}()
	_ = o.String()
	t.FailNow()
}
