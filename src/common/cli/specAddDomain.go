package cli

import (
	"strings"
)

/*
	Specification::AddDomain() implements --domain <string> flags.
*/
const (
	domainHelpText = "Specifies a fully qualified domain name."
	domainArgLong  = "domain"
)

func (o *Specification) AddDomain(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	if strings.TrimSpace(defaultValue)==""{
		panic("domain cannot be empty string.")
	}
	o.Initialize()
	o.Argument[domainArgLong] = ArgumentDescriptor{
		flagDomain,
		String,
		defaultValue,
		domainHelpText,
		ParserFlag(domainArgLong),
		ExpectNone,
	}
}
