package cli

import (
	"strconv"
)

/*
	Specification::AddConcurrency() implements --delay <int> flags.
*/
const (
	maxWordCountHelpText = "Indicates the delay each attacker will observe between attacks"
	maxWordCountArgLong  = "maxWordCount"
)

func (o *Specification) AddMaxWordCount(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[maxWordCountArgLong] = ArgumentDescriptor{
		FlagSourceMaxWordCount,
		Integer,
		strconv.Itoa(defaultValue),
		maxWordCountHelpText,
		ParserFlag(maxWordCountArgLong),
		ExpectValue,
	}
}
