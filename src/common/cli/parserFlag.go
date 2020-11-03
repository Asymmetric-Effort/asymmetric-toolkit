package cli
/*

 */
func ParserFlag(arg *string)(err error, val *Argument){
	return nil,&Argument{
		Boolean,
		"true",
	}
}
