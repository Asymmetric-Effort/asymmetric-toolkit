package source

import (
	"fmt"
)

func (o *Mode) Set(m Mode) {
	switch m {
	case Sequence, Random, Dictionary:
		*o = m
	default:
		panic(fmt.Sprintf("Invalid value in SourceMode: %d", *o))
	}
}