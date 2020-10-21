package sourcetype_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"testing"
)

func TestDataSource_Get(t *testing.T) {
	var o sourcetype.DataSource
	o = sourcetype.NotSet
	errors.Assert(o.Get() == sourcetype.NotSet, "Expected NotSet")
	o = sourcetype.Random
	errors.Assert(o.Get() == sourcetype.Random, "Expected Random")
	o = sourcetype.Sequence
	errors.Assert(o.Get() == sourcetype.Sequence, "Expected Sequence")
	o = sourcetype.Dictionary
	errors.Assert(o.Get() == sourcetype.Dictionary, "Expected Dictionary")

	//sad path
	o = BadDataSource
	defer func() { recover() }()
	_ = o.Get()
	t.FailNow()
}
