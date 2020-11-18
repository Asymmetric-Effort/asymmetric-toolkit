package cli
/*
	Specification::AddHelp() implements the -h and --help flags.
 */
const (
	helpHelpText string = "Show help / usage screen."
	helpDefault  string = "false"

	helpArgLong  = "help"
)

func (o *Specification) AddHelp() {
	o.Initialize()
	//
	// We add a long argument for help (--help)
	//
	o.Argument[helpArgLong] = ArgumentDescriptor{
		FlagHelp,
		Boolean,
		helpDefault,
		helpHelpText,
		o.ShowHelp,
		ExpectNone,
	}
}
