package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestEncryptionKey(t *testing.T) {
	var key EncryptionKey
	passphrases := []string{
		"6368616e676520746869732070617373",
		"password01",
		"cpe-1704-tks",
		"TrumpSaysMyPasswordIsSecureIfHackersDon'tGetMoreThan15Pct",
	}

	//Test the hashing...
	for i, passphrase := range passphrases {
		fmt.Println("test:", i)
		key.Set(&passphrase)
		errors.Assert(key.Length() == EncryptionKeyLength, "Expected EncryptionKeyLength")
		hash := sha256.Sum256([]byte(passphrase))
		errors.Assert(key.Equal(&hash), "Expected match.")
	}
}

func TestEncryptionKey_Equal(t *testing.T) {
	var lhs EncryptionKey
	var rhs EncryptionKey

	lhsData := "Test0"
	rhsData := "Test1"

	lhs.Set(&lhsData)
	rhs.Set(&rhsData)

	rhsBytes := [32]byte(rhs)
	errors.Assert(!lhs.Equal(&rhsBytes), "Expected inequality")
	rhsBytes = [32]byte(lhs)
	errors.Assert(lhs.Equal(&rhsBytes), "Expected equality")
}

func TestEncryptionKey_String(t *testing.T) {
	var o EncryptionKey
	passphrase:="TrumpSaysMyPasswordIsSecureIfHackersDon'tGetMoreThan15Pct"
	hash := fmt.Sprintf("%02x",sha256.Sum256([]byte(passphrase)))

	o.Set(&passphrase)
	errors.Assert(hash == o.String(), "Expected matching string output.")
}