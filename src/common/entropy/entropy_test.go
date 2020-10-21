package entropy_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)
func TestEntropy(t *testing.T) {
	errors.Assert(entropy.HighEntropyThreshold == 128, "Expected 128")
}
