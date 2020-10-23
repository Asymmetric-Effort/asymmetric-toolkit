package types

import (
	"fmt"
	"strings"
)

func (o *DataSource) Set(n string) {
	switch strings.ToLower(n) {
	case "random": *o = Random
	case "sequence": *o = Sequence
	case "dictionary": *o = Dictionary
	default:
		panic(fmt.Sprintf("Invalid DataSource value (%s)", n))
	}
}
