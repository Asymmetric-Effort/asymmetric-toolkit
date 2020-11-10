package cli

import (
	"fmt"
	"regexp"
)

const (
	sourcePatternHelpText string = "Declares a source mode (random, sequential, dictionary)."
	sourcePatternDefault  string = ".+"
	sourcePatternArgLong         = "sourcePattern"
)

func (o *Specification) AddSourcePattern(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	//
	// We add a long argument for sourcePattern (--sourcePattern)
	//
	if defaultValue == "" {
		panic("source pattern regex cannot be an empty string.")
	}
	_, err := regexp.Compile(defaultValue)
	if err != nil {
		panic(fmt.Sprintf("defaultValue must be a valid regular expression. Error:%v", err))
	}
	o.Argument[sourcePatternArgLong] = ArgumentDescriptor{
		FlagSourcePattern,
		String,
		defaultValue,
		sourcePatternHelpText,
		ParserFlag(sourcePatternArgLong),
		ExpectNone,
	}
}
