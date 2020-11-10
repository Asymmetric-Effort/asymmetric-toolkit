package cli

/*
	Specification::AddDnsRecordType() implements --dnsRecordTypes <list> flags.
*/
const (
	dnsRecordTypesHelpText = "Specifies a comma-delimited list of dns record types."
	dnsRecordTypesArgLong  = "dnsRecordTypes"
)

func (o *Specification) AddDnsRecordType(defaultValue string) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[dnsRecordTypesArgLong] = ArgumentDescriptor{
		FlagDnsRecordType,
		List,
		defaultValue,
		dnsRecordTypesHelpText,
		ParserFlag(dnsRecordTypesArgLong),
		ExpectNone,
	}
}
