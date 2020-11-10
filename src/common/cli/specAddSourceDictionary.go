package cli

/*
	Specification::AddSourceDictionary() implements --dictionary <string> flags.
*/
const (
	sourceDictionaryHelpText = "Declares a source mode (random, sequential, dictionary)"
	sourceDictionaryArgLong  = "dictionary"
	pathFileRegex = ".+" //ToDo: add regex for source path/filename.
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
		ParserString(pathFileRegex),
		ExpectValue,
	}
}
