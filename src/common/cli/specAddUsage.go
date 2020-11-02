package cli

/*
	Specification::AddUsage() will add a common flag to a given application's specification.
	By making reusable Specification methods for this purpose, we can maintain consistency
	in our toolkit CLI across different programs, reducing the learning curve.

	This specific flag adds help messaging (--help or -h) which generates usage messages based
	on a given program's specification object.

	Note: This can be overridden entirely or in part by the program's other specification definitions.
*/

const (
	helpArgLong  = "help"
	helpArgShort = "h"
)

func (o *Specification) AddUsage() {
	if _, ok := o.Argument[helpArgLong]; ok {
		o.Argument[helpArgLong] = ArgumentDescriptor{
			flagHelp,
			None,
			"",
			"Show help / usage screen.",
			o.ShowUsage,
			ExpectNone,
		}
	}
	if _, ok := o.Argument[helpArgShort]; ok {
		o.Argument[helpArgShort] = ArgumentDescriptor{
			flagHelp,
			None,
			"",
			"Show help / usage screen.",
			o.ShowUsage,
			ExpectNone,
		}
	}
}
