package entropy

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)
func TestEntropy(t *testing.T) {
	errors.Assert(HighEntropyThreshold == 128, "Expected 128")
}
