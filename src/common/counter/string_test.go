package counter

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestCounter_String(t *testing.T) {
	var c Counter
	const charset="0123456789"
	const wordSize=2
	c.Runes = func() *[]rune { d := []rune(charset); return &d }()
	c.MaxBase = uint8(len(*c.Runes)-1)
	c.Data = func() *[]uint8 { d := make([]uint8, wordSize); return &d }()
	errors.Assert(c.String()=="00", "Expect 00 in initial state.")
	(*c.Data)[0]=1
	(*c.Data)[1]=2
	errors.Assert(c.String()=="21", "Expect 21 in initial state.")
}