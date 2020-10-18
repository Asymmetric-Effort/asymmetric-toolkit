package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"crypto/sha256"
	"encoding/base64"
)

func (o *Record) Set(s string) {
	errors.Assert(o.id == nil, "Record::Set(): can only be used once to set the word value.")
	bs:=[]byte(s)
	hash := sha256.New()
	o.id = hash.Sum(bs)
	o.word = base64.StdEncoding.EncodeToString(bs)
}
