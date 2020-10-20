package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"encoding/base64"
	"testing"
)

func TestRecord_Id(t *testing.T) {
	var o Record
	testWord:= "test"
	expectedEncodedWord:=base64.StdEncoding.EncodeToString([]byte(testWord))
	errors.Assert(string(o.id) == "", "Expected empty id")
	errors.Assert(o.word == "", "Expected empty word string")
	o.id= CreateId(&testWord)
	o.word=expectedEncodedWord
	errors.Assert(o.Id()==o.id,"Test Record::Get(): expectedId.")
}