package cli

/*
	Extractor function Argument::String() will return a string value of the current argument object
	representing the internal string.  For any ArgumentType, it is possible to return the string
	value.
*/
func (o *Argument) String() string {
	return o.Value
}
