package dictionary

type Mode uint8

const (
	Noop  Mode = 00 //Error state if used.  This is the initial state.
	Read  Mode = 01 //Enable read
	Write Mode = 10 //Enable Write
)
