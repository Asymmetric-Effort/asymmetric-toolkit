package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestDataSource_IsNotSet (t *testing.T) {
	var d types.DataSource
	errors.Assert(d.IsNotSet(), "Expected NotSet")
	d.Set("Sequence")
	errors.Assert(!d.IsNotSet(), "Expected not notset")
}