package cli

/*
	Parse a no-value flag to produce a boolean value Argument.
 */

func ParserFlag() (parser func(arg *string) (err error, val *Argument)) {
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
		return nil, &Argument{
			Boolean,
			"true",
		}
	}
	//return the parser function to the Specification.
	return parser
}
