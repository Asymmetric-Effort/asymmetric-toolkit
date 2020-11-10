package cli
/*
	Specification::AddHelp() implements the -h and --help flags.
 */
const (
	usageHelpText string = "Show help / usage screen."
	usageDefault  string = ""

	helpArgLong  = "help"
	helpArgShort = "h"
)

func (o *Specification) AddHelp() {
	o.Initialize()
	//
	// We add a long argument for help (--help)
	//
	o.Argument[helpArgLong] = ArgumentDescriptor{
		FlagHelp,
		Boolean,
		usageDefault,
		usageHelpText,
		o.ShowHelp,
		ExpectNone,
	}
	//
	// We add a short argument for help (-h)
	//
	o.Argument[helpArgShort] = ArgumentDescriptor{
		FlagHelp,
		Boolean,
		usageDefault,
		usageHelpText,
		o.ShowHelp,
		ExpectNone,
	}
}
