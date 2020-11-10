package cli

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
	o.Initialize()
	o.Argument[sourceDictionaryArgLong] = ArgumentDescriptor{
		FlagSourceDictionary,
		String,
		defaultValue,
		sourceDictionaryHelpText,
		ParserFlag(sourceDictionaryArgLong),
		ExpectNone,
	}
}
