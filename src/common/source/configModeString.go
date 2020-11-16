package source

import (
	"fmt"
)

func (o *Mode) String() string{
	switch *o {
	case Sequence:
		return strModeSequence
	case Random:
		return strModeRandom
	case Dictionary:
		return strModeDictionary
	default:
		panic(fmt.Sprintf("Invalid value in SourceMode: %d", *o))
	}
}