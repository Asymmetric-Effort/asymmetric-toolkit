package counter

func (o *Counter) String() string {
	/*
		Generate a string using the current state of counter
		and the runes (character set) to which the counter will map.
	*/
	dataLen := len(*o.Data)
	var runeString []rune = make([]rune, dataLen)
	for i := uint8(0); i < uint8(dataLen); i++ {
		runeIndex := (*o.Data)[i]
		runeString[i] = (*o.Runes)[runeIndex]
	}
	return func() (result string) {
		for _, v := range runeString {
			result = string(v) + result
		}
		return
	}()
}
