package cli

import (
	"fmt"
)

/*
	ParserEnum() is a parser function factory which returns a parser function to the caller.
	This parser function will evaluate a given commandline argument (arg *string) and perform
	any validation before returning an appropriate error and Argument object pointer.
*/

func ParserEnum(p ...string) (parser func(arg *string) (err error, val *Argument)) {
	/*
		Create the parser function and return it at the end.

		This Parser Factory function will take a list of parameters (p) where each parameter
		represents an item in the enumerated set.  We will then return a parser function which
		has knowledge of this set to ensure any commandline argument string is an element in the
		aforesaid set and if so return an Argument object.
	*/
	var enumeratedSet []string // The enumerated set we will bounds check for.

	if len(p) == 0 {
		//
		// Ensure we do not have an empty set which would be meaningless.
		//
		panic("ParserEnum() cannot create a parser function for an empty set.")
	}
	for _, e := range p {
		enumeratedSet = append(enumeratedSet, e)
	}

	find := func(s *string) bool {
		for _, element := range enumeratedSet {
			if element == *s {
				return true
			}
		}
		return false
	}

	parser = func(arg *string) (err error, val *Argument) {
		//
		// When our parser function runs, we validate whether the given commandline argument
		// matches parameter passed to the parser function in the Specification.
		//
		if find(arg) {
			return nil, &Argument{
				String,
				*arg,
			}
		} else {
			//
			// Bound check failed.  Return an error and nil Argument object.
			//
			return fmt.Errorf("unexpected value (%s) not in set (%v)",*arg,enumeratedSet), nil
		}
	}
	//return the parser function to the Specification.
	return parser
}
