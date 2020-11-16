package counter
/*
	Counter struct is the data structure used to store the internal state of the counter.
 */
type Counter struct {
	runes *[]rune	//Allowed chars in the character set.
	data *[]uint8
	maxBase uint8
}
