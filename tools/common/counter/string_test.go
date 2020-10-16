package counter

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestCounter_String(t *testing.T) {
	var c Counter
	const charset="0123456789"
	const wordSize=2
	c.runes = func() *[]rune { d := []rune(charset); return &d }()
	c.maxBase = uint8(len(*c.runes)-1)
	c.data = func() *[]uint8 { d := make([]uint8, wordSize); return &d }()
	errors.Assert(c.String()=="00", "Expect 00 in initial state.")
	(*c.data)[0]=1
	(*c.data)[1]=2
	errors.Assert(c.String()=="21", "Expect 21 in initial state.")
}