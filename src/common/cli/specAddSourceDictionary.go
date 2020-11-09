package cli

import (
	"strings"
)

/*
	Specification::AddSourceDictionary() implements --dictionary <string> flags.
*/
const (
	sourceDictionaryHelpText = "Declares a source mode (random, sequential, dictionary)"
	sourceDictionaryArgLong  = "dictionary"
)

func (o *Specification) AddSourceDictionary(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	if strings.TrimSpace(defaultValue) == "" {
		panic("source file/path cannot be empty string.")
	}
	o.Initialize()
	o.Argument[sourceDictionaryArgLong] = ArgumentDescriptor{
		flagSourceDictionary,
		String,
		defaultValue,
		sourceDictionaryHelpText,
		ParserFlag(sourceDictionaryArgLong),
		ExpectNone,
	}
}
