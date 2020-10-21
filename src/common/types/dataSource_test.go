package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

const (
	BadDataSource = 255
)

func TestDataSourceInitialState(t *testing.T){
	var d DataSource
	errors.Assert(d == NotSet, "Expected initial Value to be NotSet")
}

func TestDataSourceValues(t *testing.T){
	errors.Assert(NotSet == 0, "Expected 0")
	errors.Assert(Sequence == 1, "Expected 1")
	errors.Assert(Random == 2, "Expected 2")
	errors.Assert(Dictionary == 3, "Expected 3")
}