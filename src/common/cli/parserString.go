package cli

import (
	"fmt"
	"regexp"
)

/*
	ParserString() is a parser function factory which returns a parser function to the caller.
	This parser function will evaluate a given commandline argument (arg *string) and perform
	any validation before returning an appropriate error and Argument object pointer.
*/

func ParserString(p string) (parser func(arg *string) (err error, val *Argument)) {
	/*
		Create the parser function and return it at the end.

		The parser will be executed by CommandLine::Parser() during runtime, but
		this factory function will be executed by the Specification to create the
		parser function.

		This string will be vetted against a pattern using a regular expression passed to the factory
		function.
	*/
	if p == "" {
		p = ".*" // By default we will accept any string (.*)
	}
	re := regexp.MustCompile(p)
	parser = func(arg *string) (err error, val *Argument) {
		//
		// When our parser function runs, we validate whether the given commandline argument
		// matches the parameter passed to the parser function in the Specification.
		//
		if re.MatchString(*arg) {
			//
			// When the string matches, we return an Argument object, properly configured.
			// and no error (nil) is returned.
			//
			return nil, &Argument{
				String,
				*arg,
			}
		} else {
			//
			// If a commandline argument string is evaluated against the regular expression,
			// and no match is found, we will return an error object and a nil Argument pointer.
			//
			return fmt.Errorf("error: String fails to match " +
				"pattern (%s) when evaluating '%s'", p, *arg), nil
		}
	}
	//return the parser function to the Specification.
	return parser
}

