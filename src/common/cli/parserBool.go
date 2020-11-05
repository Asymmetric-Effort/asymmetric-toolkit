package cli

/*
	ParserBool() is a parser function factory which returns a parser function to the caller.
	This parser function will evaluate a given commandline argument (arg *string) and perform
	any validation before returning an appropriate error and Argument object pointer.
*/

import (
	"fmt"
	"strconv"
)

func ParserBool() (parser func(arg *string) (err error, val *Argument)) {
	/*
		Create the parser function and return it at the end.

		The parser will be executed by CommandLine::Parser() during runtime, but
		this factory function will be executed by the Specification to create the
		parser function.
	*/
	parser = func(arg *string) (err error, val *Argument) {
		//
		// When our parser function runs, we trump and shift the command line argument to lower-case
		// so our evaluation is case insensitive.  This allows a simple evaluation and validation.
		//
		if _, e := strconv.ParseBool(*arg); e != nil {
			err = fmt.Errorf("error parsing expected boolean value.  Encountered (%s): %v", *arg, e)
			val = nil
		} else {
			err = nil
			val = &Argument{
				Boolean,
				*arg,
			}
		}
		return
	}
	//return the parser function to the Specification.
	return parser
}
