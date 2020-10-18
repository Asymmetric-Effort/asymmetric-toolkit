package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRecord_Set(t *testing.T) {
	var o Record
	const testWord = "test"
	expectedEncodedWord:=base64.StdEncoding.EncodeToString([]byte(testWord))
	hash:=sha256.New()
	expectedId := hash.Sum([]byte(testWord))
	errors.Assert(string(o.id) == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.Set(testWord)
	errors.Assert(bytes.Equal(o.id, expectedId), fmt.Sprintf("Expected empty id:\n%x\n%x", o.id, expectedId))
	errors.Assert(o.word == expectedEncodedWord, fmt.Sprintf("Expected empty word string: %s", o.word))
}
