package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"encoding/base64"
)

func (o *Record) Get() string {
	w, err := base64.StdEncoding.DecodeString(o.word)
	errors.Assert(err == nil, "Record::Get() an error occurred getting the word content of a dictionary")
	return string(w)
}