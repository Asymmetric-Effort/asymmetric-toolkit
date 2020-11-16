package counter

/*
	Counter::Setup() initializes the counter and verifies the inputs which will be used
	by the counter to generate a sequential set of words using a given character set.
*/
import "asymmetric-effort/asymmetric-toolkit/src/common/errors"

func (o *Counter) Setup(charset *string, wordSize int) {
	//
	// Validate the inputs.
	//
	errors.Assert(*charset != "", "Counter::Setup() expects non empty string for charset")
	errors.Assert(wordSize > 0, "Counter::Setup() expects positive integer word size.")
	//
	// Store the character set as an array of runes.
	//
	o.runes = func() *[]rune { d := []rune(*charset); return &d }()
	//
	// Create the base array
	//
	o.maxBase = uint8(len(*o.runes) - 1)
	//
	// Create the data array used for the counter.  this array provides what amounts to wheels in
	// an odometer which roll over in sequence, counting through the character set in ascending order.
	//
	o.data = func() *[]uint8 { d := make([]uint8, wordSize); return &d }()
}
