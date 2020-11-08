package cli

import (
	"strings"
)

/*
	ParseList() is a parser function factory which returns a parser function to the caller.
	This parser function will evaluate a given commandline argument (arg *string) and perform
	any validation before returning an appropriate error and Argument object pointer.
*/

func ParseList(delimiter string) (parser func(arg *string) (err error, val *Argument)) {
	/*
		Create the parser function and return it at the end.
	*/
	parser = func(arg *string) (err error, val *Argument) {
		//
		//	Start off by converting our string commandline argument to its numeric equivalent.
		//
		_ = strings.Split(*arg, delimiter)
		return nil, &Argument{
			List,
			*arg,
		}
	}
	//return the parser function to the Specification.
	return parser
}
