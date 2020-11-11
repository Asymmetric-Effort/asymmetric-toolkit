package cli

/*
	Specification::AddSourceDictionary() implements --dictionary <string> flags.
*/

import (
	"fmt"
)

const (
	sourceDictionaryHelpText = "Declares a source mode (random, sequential, dictionary)"
	sourceDictionaryArgLong  = "dictionary"
	pathFileRegex            = ".+" //ToDo: add regex for source path/filename.
)

func (o *Specification) AddSourceDictionary(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	if defaultValue == "" {
		panic("Dictionary file default value cannot be an empty string.")
	}
	o.Initialize()
	o.Argument[sourceDictionaryArgLong] = ArgumentDescriptor{
		FlagSourceDictionary,
		String,
		defaultValue,
		sourceDictionaryHelpText,
		ParserString(pathFileRegex),
		ExpectValue,
	}
	fmt.Println("Specification::AddSourceDictionary()")
}
