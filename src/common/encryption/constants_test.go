package encryption_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestEncryptionConstant(t *testing.T) {
	errors.Assert(encryption.NonceSize == 12, "Expected nonce size of 12")
}
