package types

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestDataSource_Get(t *testing.T) {
	var o DataSource
	o = NotSet
	errors.Assert(o.Get() == NotSet, "Expected NotSet")
	o = Random
	errors.Assert(o.Get() == Random, "Expected Random")
	o = Sequence
	errors.Assert(o.Get() == Sequence, "Expected Sequence")
	o = Dictionary
	errors.Assert(o.Get() == Dictionary, "Expected Dictionary")
}

func TestDataSource_Get_Bad(t *testing.T) {
	var o DataSource = BadDataSource
	defer func() { recover() }()
	_ = o.Get()
	t.FailNow()
}
