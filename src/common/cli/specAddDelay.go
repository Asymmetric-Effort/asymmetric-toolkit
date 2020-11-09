package cli

import (
	"strconv"
)

/*
	Specification::AddConcurrency() implements --delay <int> flags.
*/
const (
	delayHelpText = "Indicates the delay each attacker will observe between attacks"
	delayArgLong  = "delay"
)

func (o *Specification) AddDelay(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[delayArgLong] = ArgumentDescriptor{
		flagDelay,
		Integer,
		strconv.Itoa(defaultValue),
		delayHelpText,
		ParserFlag(delayArgLong),
		ExpectNone,
	}
}
