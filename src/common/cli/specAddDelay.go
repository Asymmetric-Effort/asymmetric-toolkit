package cli

import (
	"math"
	"strconv"
)

/*
	Specification::AddDelay() implements --delay <int> flags.
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
		FlagDelay,
		Integer,
		strconv.Itoa(defaultValue),
		delayHelpText,
		ParserInt(0, math.MaxInt32),
		ExpectValue,
	}
}
