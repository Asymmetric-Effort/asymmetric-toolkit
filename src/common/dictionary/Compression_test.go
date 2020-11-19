package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDictionary_CompressionType(t *testing.T) {
	var o CompressionAlgorithms
	errors.Assert(o == notImplemented, "expect 0")
	errors.Assert(notImplemented == 0, "Expect 0")
	errors.Assert(gzip == 1, "expect 1")
}
