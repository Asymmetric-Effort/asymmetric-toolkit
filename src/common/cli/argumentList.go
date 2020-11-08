package cli

import (
	"strconv"
	"strings"
)

/*
	Extractor function Argument::List() will convert the internal (string) state to an array of strings
	split with the Parser defined delimiter.
*/

func (o *Argument) List() []string {
	/*
		Extract and return our list of strings.
	*/
	getDelimiter := func() string {
		//
		// Get the delimiter from our encoded list string (delimiterLen,delimiterStr,ListString)
		//
		delimLength, err := strconv.Atoi(string(o.Value))
		if err != nil {
			panic(err)
		}
		//
		// Extract the delimiter and return the same.
		//
		return o.Value[1 : delimLength+1] //get the delimiter
	}

	if o.Type == List {
		//
		// extract and split our list using the encoded delimiter.
		//
		return strings.Split(o.Value[2:], getDelimiter())
	} else {
		panic("Type mismatch.  Attempted to extract float from non-float Argument.")
	}
}
