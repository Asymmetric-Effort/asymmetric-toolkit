package cli

import (
	"strings"
)

/*
	Specification::AddDomain() implements --domain <string> flags.
*/
const (
	outputHelpText = "Specifies a fully qualified domain name."
	outputArgLong  = "output"
)

func (o *Specification) AddOutput(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	if strings.TrimSpace(defaultValue)==""{
		panic("output cannot be empty string.")
	}
	o.Initialize()
	o.Argument[outputArgLong] = ArgumentDescriptor{
		FlagOutput,
		String,
		defaultValue,
		outputHelpText,
		ParserFlag(domainArgLong),
		ExpectNone,
	}
}
