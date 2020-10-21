package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestDataSource_IsRandom (t *testing.T) {
	var d types.DataSource
	d.Set("Random")
	errors.Assert(d.IsRandom(), "Expected random(1)")
	d.Set("random")
	errors.Assert(d.IsRandom(), "Expected random(2)")
	d.Set("Sequence")
	errors.Assert(!d.IsRandom(), "Expected not random")
}