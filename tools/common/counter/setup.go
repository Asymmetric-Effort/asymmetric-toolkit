package counter

func (o *Counter) Setup(charset string, wordSize int) {
	o.runes = func() *[]rune { d := []rune(charset); return &d }()
	o.maxBase = uint8(len(*o.runes))
	o.data = func() *[]uint8 { d := make([]uint8, wordSize); return &d }()
}
