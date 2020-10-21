package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDataSource_IsDictionary (t *testing.T) {
	var d DataSource
	d.Set("Dictionary")
	errors.Assert(d.IsDictionary(), "Expected Dictionary(1)")
	d.Set("dictionary")
	errors.Assert(d.IsDictionary(), "Expected Dictionary(2)")
	d.Set("Sequence")
	errors.Assert(!d.IsDictionary(), "Expected not dictionary")
}