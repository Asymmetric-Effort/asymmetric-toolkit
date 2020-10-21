package sourcetype_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"strconv"
	"testing"
)

func TestDataSource_Set(t *testing.T) {
	var o sourcetype.DataSource
	o.Set("Random")
	errors.Assert(o == sourcetype.Random, "Expected Random")
	o.Set("random")
	errors.Assert(o == sourcetype.Random, "Expected Random")
	o.Set("Sequence")
	errors.Assert(o == sourcetype.Sequence, "Expected Sequence")
	o.Set("sequence")
	errors.Assert(o == sourcetype.Sequence, "Expected Sequence")
	o.Set("Dictionary")
	errors.Assert(o == sourcetype.Dictionary, "Expected Dictionary")
	o.Set("dictionary")
	errors.Assert(o == sourcetype.Dictionary, "Expected Dictionary")

	//Sad Path
	defer func() { recover() }()
	o.Set(strconv.Itoa(int(BadDataSource)))
	t.FailNow()
}
