package cli

const (
	debugHelpText string = "Show debug logging"
	debugDefault  string = ""
	debugArgLong         = "debug"
)

func (o *Specification) AddDebug() {
	//
	// Initialize the Argument object.
	//
	if o.Argument == nil {
		o.Argument = make(map[string]ArgumentDescriptor)
		o.Argument[""] = ArgumentDescriptor{}
	}
	//
	// We add a long argument for debug (--debug)
	//
	o.Argument[debugArgLong] = ArgumentDescriptor{
		flagDebug,
		Boolean,
		debugDefault,
		debugHelpText,
		ParserFlag,
		ExpectNone,
	}
}
