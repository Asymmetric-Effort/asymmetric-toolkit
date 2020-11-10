package cli

import (
	"math"
	"strconv"
)

/*
	Specification::AddDepth() implements --depth <int> flags.
*/
const (
	depthHelpText = "Indicates the number of subdomain levels into a DNS domain to enumerate."
	depthArgLong  = "depth"
)

func (o *Specification) AddDepth(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[depthArgLong] = ArgumentDescriptor{
		FlagDepth,
		Integer,
		strconv.Itoa(defaultValue),
		depthHelpText,
		ParserInt(0, math.MaxInt16),
		ExpectValue,
	}
}
