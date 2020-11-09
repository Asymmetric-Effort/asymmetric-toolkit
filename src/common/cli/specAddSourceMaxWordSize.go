package cli

import (
	"strconv"
)

/*
	Specification::AddConcurrency() implements --delay <int> flags.
*/
const (
	maxWordSizeHelpText = "Defines the maximum word size a generator should target."
	maxWordSizeArgLong  = "maxWordSize"
)

func (o *Specification) AddMaxWordSize(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[maxWordSizeArgLong] = ArgumentDescriptor{
		flagSourceMaxWordSize,
		Integer,
		strconv.Itoa(defaultValue),
		maxWordSizeHelpText,
		ParserFlag(maxWordCountArgLong),
		ExpectNone,
	}
}
