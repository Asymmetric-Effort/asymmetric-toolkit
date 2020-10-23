package sourcetype_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source/sourcetype"
	"testing"
)

const (
	BadDataSource = 255
)

func TestConstants(t *testing.T) {
	var d sourcetype.DataSource
	errors.Assert(d == sourcetype.NotSet, "Expected initial Value to be NotSet")
	errors.Assert(sourcetype.NotSet == 0, "Expected 0")
	errors.Assert(sourcetype.Sequence == 1, "Expected 1")
	errors.Assert(sourcetype.Random == 2, "Expected 2")
	errors.Assert(sourcetype.Dictionary == 3, "Expected 3")
}
