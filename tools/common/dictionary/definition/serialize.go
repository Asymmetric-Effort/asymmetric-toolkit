package dictionaryDefinition

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func (o *Record) Serialize() string {
	//ToDo: Add digested stats after hash (id) and before wordLen


	return strings.Join([]string{
		string(o.id),
		strconv.Itoa(len(o.word)),/* word length provides safety for reading the word later.*/
		base64.StdEncoding.EncodeToString([]byte(o.word))},
		delimiter)
}
