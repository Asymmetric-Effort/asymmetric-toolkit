package cli

import (
	"strings"
)

/*
	Specification::AddDnsServer() implements --dnsServer <string> flags.
*/
const (
	dnsServerHelpText = "Specifies the DNS server to be used by attackers."
	dnsServerArgLong  = "dnsServer"
	fqdnOrIpRegex    = ".+"
	//ToDo: Add regex string pattern for dns fqdn or ip address.
)

func (o *Specification) AddDnsServer(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	if strings.TrimSpace(defaultValue) == "" {
		panic("dnsServer cannot be empty string.")
	}
	o.Initialize()
	o.Argument[dnsServerArgLong] = ArgumentDescriptor{
		FlagDnsServer,
		String,
		defaultValue,
		dnsServerHelpText,
		ParserString(fqdnOrIpRegex),
		ExpectValue,
	}
}
