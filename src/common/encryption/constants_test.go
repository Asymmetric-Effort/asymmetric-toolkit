package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestEncryptionConstant(t *testing.T) {
	errors.Assert(nonceSize == 12, "Expected nonce size of 12")
}
