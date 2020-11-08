package cli

import (
	"fmt"
	"strings"
)

/*
	ParserList() is a parser function factory which returns a parser function to the caller.
	This parser function will evaluate a given commandline argument (arg *string) and perform
	any validation before returning an appropriate error and Argument object pointer.
*/

func ParserList(delimiter string) (parser func(arg *string) (err error, val *Argument)) {
	/*
		Create the parser function and return it at the end.
	*/
	parser = func(arg *string) (err error, val *Argument) {
		//
		// Split our string with our delimiter.  This just confirms there are no errors.
		//
		_ = strings.Split(*arg, delimiter)
		//
		// Encode our list and return our Argument pointer.
		// Note that we prepend the list string with the delimiter length and delimiter string.
		// So that the Argument::List() method can extract the list on request.
		//
		return nil, &Argument{
			List,
			fmt.Sprintf("%d%s%s", len(delimiter), delimiter, *arg),
		}
	}
	//return the parser function to the Specification.
	return parser
}
