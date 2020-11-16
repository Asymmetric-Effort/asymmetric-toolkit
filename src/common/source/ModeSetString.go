package source

import (
	"fmt"
	"strings"
)

func (o *Mode) SetString(s string) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case strModeSequence:
		*o = Sequence
	case strModeRandom:
		*o = Random
	case strModeDictionary:
		*o = Dictionary
	default:
		panic(fmt.Sprintf("Invalid string value passed to SetString(): %s", s))
	}
}
