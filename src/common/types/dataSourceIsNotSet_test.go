package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDataSource_IsNotSet (t *testing.T) {
	var d DataSource
	errors.Assert(d.IsNotSet(), "Expected NotSet")
	d.Set("Sequence")
	errors.Assert(!d.IsNotSet(), "Expected not notset")
}