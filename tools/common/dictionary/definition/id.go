package dictionaryDefinition

import (
	"fmt"
)

func (o *Record) Id() string {
	return fmt.Sprintf("%x",o.id)
}
