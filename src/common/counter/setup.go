package counter

import "asymmetric-effort/asymmetric-toolkit/src/common/errors"

func (o *Counter) Setup(charset string, wordSize int) {
	errors.Assert(charset != "", "Counter::Setup() expects non empty string for charset")
	errors.Assert(wordSize > 0, "Counter::Setup() expects positive integer word size.")
	o.Runes = func() *[]rune { d := []rune(charset); return &d }()
	o.MaxBase = uint8(len(*o.Runes)-1)
	o.Data = func() *[]uint8 { d := make([]uint8, wordSize); return &d }()
}
