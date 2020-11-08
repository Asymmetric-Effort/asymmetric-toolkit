package cli

const (
	usageHelpText string = "Show help / usage screen."
	usageDefault  string = ""

	helpArgLong  = "help"
	helpArgShort = "h"
)

func (o *Specification) AddUsage() {
	//
	// Initialize the Argument object.
	//
	if o.Argument == nil {
		o.Argument = make(map[string]ArgumentDescriptor)
		o.Argument[""] = ArgumentDescriptor{}
	}
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
