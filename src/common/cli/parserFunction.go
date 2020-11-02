package cli
/*
	Our Parser Function type represents the signature of all parser functions.  For every argument a program processes,
	it should provide a context-specific Parse() function which will take a string value (argument) and return the
	pointer to an Argument containing the processed value.

	If the ParserFunction pointer is nil, then the CommandLine::Parse() function will assume a noop.
	Programs might be more specific and use ParserNoop() instead as this will allow future logging facilities, etc.
 */
type ParserFunction func() (err error, val *Argument)
