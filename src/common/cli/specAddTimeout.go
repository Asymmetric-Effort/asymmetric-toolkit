package cli

import (
	"math"
	"strconv"
)

/*
	Specification::AddTimeout() implements --timeout <int> flags.
*/
const (
	timeoutHelpText = "Indicates the connection timeout (in seconds).  Must be positive integer."
	timeoutArgLong  = "timeout"
	timeoutDefault  = 60
)

func (o *Specification) AddTimeout(defaultValue int) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[timeoutArgLong] = ArgumentDescriptor{
		FlagTimeout,
		Integer,
		strconv.Itoa(func()int{
			if defaultValue == 0 {
				return timeoutDefault
			}else{
				return defaultValue
			}
		}()),
		timeoutHelpText,
		ParserInt(0,math.MaxInt32),
		ExpectNone,
	}
}
