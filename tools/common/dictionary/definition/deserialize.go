package dictionaryDefinition

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"encoding/base64"
	"strconv"
	"strings"
)

func (o *Record) Deserialize(s string) {
	const (
		fieldId      = 0
		fieldWordLen = 1
		fieldWord    = 2
	)
	fields := strings.Split(s, delimiter)
	errors.Assert(len(fields) == 3, "dictionaryDefinition/Record::Deserialize(): ")
	o.id = []byte(fields[fieldId])
	wordLen := func() int {
		i, err := strconv.Atoi(fields[fieldWordLen])
		errors.Assert(err == nil, "Error parsing word length field from dictionary definition record.")
		return i
	}()
	o.word = func() string {
		s,err:=base64.StdEncoding.DecodeString(fields[fieldWord])
		errors.Assert(err==nil, "Error decoding word string from dictionary definition record.")
		return string(s)
	}()
	errors.Assert(wordLen == len(o.word), "Word length mismatch")
}
