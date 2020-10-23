package reader_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/reader"
	"testing"
)

func TestReader_IsNil(t *testing.T) {
	var r reader.Reader
	errors.Assert(r.IsNil(), "Expected nil initial state")
	t.SkipNow()
}
