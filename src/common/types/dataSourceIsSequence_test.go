package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestDataSource_IsSequence (t *testing.T) {
	var d types.DataSource
	d.Set("Sequence")
	errors.Assert(d.IsSequence(), "Expected sequence(1)")
	d.Set("sequence")
	errors.Assert(d.IsSequence(), "Expected sequence(2)")
	d.Set("Random")
	errors.Assert(!d.IsSequence(), "Expected not sequence")
}