package cli
/*
	Parse a no-value flag to produce a boolean value Argument.
 */
func ParserFlag(arg *string)(err error, val *Argument){
	return nil,&Argument{
		Boolean,
		"true",
	}
}
