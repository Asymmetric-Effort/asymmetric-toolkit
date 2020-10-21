package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

const (
	BadDataSource = 255
)

func TestDataSourceInitialState(t *testing.T){
	var d types.DataSource
	errors.Assert(d == types.NotSet, "Expected initial Value to be NotSet")
}

func TestDataSourceValues(t *testing.T){
	errors.Assert(types.NotSet == 0, "Expected 0")
	errors.Assert(types.Sequence == 1, "Expected 1")
	errors.Assert(types.Random == 2, "Expected 2")
	errors.Assert(types.Dictionary == 3, "Expected 3")
}