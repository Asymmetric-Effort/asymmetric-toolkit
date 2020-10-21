package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestDataSource_Get(t *testing.T) {
	var o types.DataSource
	o = types.NotSet
	errors.Assert(o.Get() == types.NotSet, "Expected NotSet")
	o = types.Random
	errors.Assert(o.Get() == types.Random, "Expected Random")
	o = types.Sequence
	errors.Assert(o.Get() == types.Sequence, "Expected Sequence")
	o = types.Dictionary
	errors.Assert(o.Get() == types.Dictionary, "Expected Dictionary")
}

func TestDataSource_Get_Bad(t *testing.T) {
	var o types.DataSource = BadDataSource
	defer func() { recover() }()
	_ = o.Get()
	t.FailNow()
}
