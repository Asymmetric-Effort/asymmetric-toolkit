package cli

import (
	"strconv"
)

/*
	Specification::AddMinWordSize() implements --minWordSize <int> flags.
*/
const (
	minWordSizeHelpText = "Defines the minimum word size a generator should target."
	minWordSizeArgLong  = "minWordSize"
)

func (o *Specification) AddMinWordSize(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[minWordSizeArgLong] = ArgumentDescriptor{
		FlagSourceMinWordSize,
		Integer,
		strconv.Itoa(defaultValue),
		minWordSizeHelpText,
		ParserFlag(minWordSizeArgLong),
		ExpectNone,
	}
}
