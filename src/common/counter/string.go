package counter

func (o *Counter) String() string {
	/*
		Generate a string using the current state of counter
		and the runes (character set) to which the counter will map.
	*/
	if (o.runes == nil) || (o.data == nil) {
		// Make sure we don't try to operate on nil data fields.
		return ""
	}
	dataLen := len(*o.data)
	var runeString []rune = make([]rune, dataLen)
	for i := uint8(0); i < uint8(dataLen); i++ {
		runeIndex := (*o.data)[i]
		runeString[i] = (*o.runes)[runeIndex]
	}
	return func() (result string) {
		for _, v := range runeString {
			result = string(v) + result
		}
		return
	}()
}
