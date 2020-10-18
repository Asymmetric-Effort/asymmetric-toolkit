package definition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"crypto/sha256"
	"encoding/base64"
)

type Record struct {
	id      []byte
	word    string //Base64-encoded word
}

func (o *Record) Set(s string) {
	errors.Assert(o.id == nil, "Record::Set(): can only be used once to set the word value.")
	bs:=[]byte(s)
	hash := sha256.New()
	o.id = hash.Sum(bs)
	o.word = base64.StdEncoding.EncodeToString(bs)
}

func (o *Record) Get() string {
	w, err := base64.StdEncoding.DecodeString(o.word)
	errors.Assert(err == nil, "Record::Get() an error occurred getting the word content of a dictionary")
	return string(w)
}

func (o *Record) Id() string {
	return string(o.id)
}
