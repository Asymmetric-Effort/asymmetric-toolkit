package definition_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/encryption"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/definition"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestRecord_Get(t *testing.T) {
	var o definition.Record
	var key encryption.Key

	testWord := "test"
	passphrase := "myPassphrase"

	key.Set(&passphrase)

	expectedId := fmt.Sprintf("%x", sha256.Sum256([]byte(testWord)))
	errors.Assert(o.ID() == "", "Expected empty id")
	errors.Assert(o.String(&key) == "", "Expected empty word string")
	o.Set(testWord, &key)
	errors.Assert(o.Get(&key) == testWord, "Test Record::Get(): Expected test word mismatch.")
	errors.Assert(o.ID() == expectedId, "Expected sha256 hash")
}
