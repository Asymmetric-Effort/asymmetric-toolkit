package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

func TestRecord_Get(t *testing.T) {
	var o Record
	const testWord = "test"
	expectedEncodedWord:=base64.StdEncoding.EncodeToString([]byte(testWord))
	hash:=sha256.New()
	expectedId := hash.Sum([]byte(testWord))
	errors.Assert(string(o.id) == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.id=expectedId
	o.word=expectedEncodedWord
	errors.Assert(o.Get()==testWord,"Test Record::Get(): Expected test word mismatch.")

}
