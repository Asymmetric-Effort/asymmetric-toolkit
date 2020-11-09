package cli
/*
	Specification::AddUsage() implements the -h and --help flags.
 */
const (
	usageHelpText string = "Show help / usage screen."
	usageDefault  string = ""

	helpArgLong  = "help"
	helpArgShort = "h"
)

func (o *Specification) AddUsage() {
	o.Initialize()
	//
	// We add a long argument for help (--help)
	//
	o.Argument[helpArgLong] = ArgumentDescriptor{
		flagHelp,
		None,
		usageDefault,
		usageHelpText,
		o.ShowUsage,
		ExpectNone,
	}
	//
	// We add a short argument for help (-h)
	//
	o.Argument[helpArgShort] = ArgumentDescriptor{
		flagHelp,
		None,
		usageDefault,
		usageHelpText,
		o.ShowUsage,
		ExpectNone,
	}
}
