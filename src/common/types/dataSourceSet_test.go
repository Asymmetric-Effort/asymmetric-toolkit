package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"strconv"
	"testing"
)

func TestDataSource_Set(t *testing.T) {
	var o DataSource
	o.Set("Random")
	errors.Assert(o == Random, "Expected Random")
	o.Set("random")
	errors.Assert(o == Random, "Expected Random")
	o.Set("Sequence")
	errors.Assert(o == Sequence, "Expected Sequence")
	o.Set("sequence")
	errors.Assert(o == Sequence, "Expected Sequence")
	o.Set("Dictionary")
	errors.Assert(o == Dictionary, "Expected Dictionary")
	o.Set("dictionary")
	errors.Assert(o == Dictionary, "Expected Dictionary")
}

func TestDataSource_Set_Bad(t *testing.T) {
	var o DataSource
	defer func(){recover()}()
	o.Set(strconv.Itoa(int(BadDataSource)))
	t.FailNow()
}
