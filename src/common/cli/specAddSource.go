package cli

import (
	"strings"
)

/*
	Specification::AddSource() implements --source <string> flags.
*/
const (
	sourceHelpText = "Declares a source mode (random, sequential, dictionary)."
	sourceArgLong  = "source"
)

func (o *Specification) AddSource(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	if strings.TrimSpace(defaultValue)==""{
		panic("domain cannot be empty string.")
	}
	o.Initialize()
	o.Argument[sourceArgLong] = ArgumentDescriptor{
		FlagSource,
		String,
		defaultValue,
		sourceHelpText,
		ParserFlag(sourceArgLong),
		ExpectNone,
	}
}
