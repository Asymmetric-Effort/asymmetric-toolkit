package cli

import (
	"strconv"
	"strings"
)

/*
	Extractor function Argument::String() will return a string value of the current argument object
	representing the internal string.  For any ArgumentType, it is possible to return the string
	value.
*/
func (o *Argument) String() string {
	getDelimiter := func() string {
		//
		// Get the delimiter from our encoded list string (delimiterLen,delimiterStr,ListString)
		//
		delimLength, err := strconv.Atoi(string(o.Value[0]))
		if err != nil {
			panic(err)
		}
		//
		// Extract the delimiter and return the same.
		//
		return o.Value[1 : delimLength+1] //get the delimiter
	}
	switch o.Type {
	case List:
		return strings.Join(o.List(),getDelimiter())
	default:
		return o.Value
	}
}
