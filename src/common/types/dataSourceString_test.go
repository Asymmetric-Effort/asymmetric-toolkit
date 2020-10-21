package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestDataSource_String(t *testing.T) {
	var o types.DataSource
	o = types.Random
	errors.Assert(o.String() == "Random", "Expected Random")
	o = types.Sequence
	errors.Assert(o.String() == "Sequence", "Expected Sequence")
	o = types.Dictionary
	errors.Assert(o.String() == "Dictionary", "Expected Dictionary")
}

func TestDataSource_String_Bad(t *testing.T) {
	var o types.DataSource = BadDataSource
	defer func(){recover()}()
	_ = o.String()
	t.FailNow()
}
