package types

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestDataSource_IsNotSet (t *testing.T) {
	var d DataSource
	errors.Assert(d.IsNotSet(), "Expected NotSet")
	d.Set("Sequence")
	errors.Assert(!d.IsNotSet(), "Expected not notset")
}