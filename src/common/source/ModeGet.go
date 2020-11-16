package source

import (
	"fmt"
)

func (o *Mode) Get() Mode {
	switch *o {
	case Sequence, Random, Dictionary:
		return *o
	default:
		panic(fmt.Sprintf("Invalid value in SourceMode: %d", *o))
	}
}
