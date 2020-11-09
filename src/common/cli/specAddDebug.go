package cli

const (
	debugHelpText string = "Show debug logging"
	debugDefault  string = "false"
	debugArgLong         = "debug"
)

func (o *Specification) AddDebug() {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	//
	// We add a long argument for debug (--debug)
	//
	o.Argument[debugArgLong] = ArgumentDescriptor{
		flagDebug,
		Boolean,
		debugDefault,
		debugHelpText,
		ParserFlag(debugArgLong),
		ExpectNone,
	}
}
