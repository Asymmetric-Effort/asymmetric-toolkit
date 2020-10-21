package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"strconv"
	"testing"
)

func TestDataSource_Set(t *testing.T) {
	var o types.DataSource
	o.Set("Random")
	errors.Assert(o == types.Random, "Expected Random")
	o.Set("random")
	errors.Assert(o == types.Random, "Expected Random")
	o.Set("Sequence")
	errors.Assert(o == types.Sequence, "Expected Sequence")
	o.Set("sequence")
	errors.Assert(o == types.Sequence, "Expected Sequence")
	o.Set("Dictionary")
	errors.Assert(o == types.Dictionary, "Expected Dictionary")
	o.Set("dictionary")
	errors.Assert(o == types.Dictionary, "Expected Dictionary")
}

func TestDataSource_Set_Bad(t *testing.T) {
	var o types.DataSource
	defer func(){recover()}()
	o.Set(strconv.Itoa(int(BadDataSource)))
	t.FailNow()
}
