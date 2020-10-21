package encryption_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestEncryptionConstant(t *testing.T) {
	errors.Assert(encryption.NonceSize == 12, "Expected nonce size of 12")
	errors.Assert(encryption.KeyLength == 32, "Expected Key length to be 32 bytes (sha256)")
}
