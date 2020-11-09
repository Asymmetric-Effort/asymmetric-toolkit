package cli


const (
	sourcePatternHelpText string = "Show sourcePattern logging"
	sourcePatternDefault  string = ".+"
	sourcePatternArgLong         = "sourcePattern"
)

func (o *Specification) AddSourcePattern() {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	//
	// We add a long argument for sourcePattern (--sourcePattern)
	//
	o.Argument[sourcePatternArgLong] = ArgumentDescriptor{
		flagSourcePattern,
		Boolean,
		sourcePatternDefault,
		sourcePatternHelpText,
		ParserFlag(sourcePatternArgLong),
		ExpectNone,
	}
}
