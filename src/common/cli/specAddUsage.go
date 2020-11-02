package cli

func (o *Specification) AddUsage() {
	//ToDo: Do not overwrite if it exists.
	o.Argument[helpArgLong] = ArgumentDescriptor{
		flagHelp,
		None,
		"",
		"Show help / usage screen.",
		o.ShowUsage,
		ExpectEnd,
	}
	o.Argument[helpArgShort] = ArgumentDescriptor{
		flagHelp,
		None,
		"",
		"Show help / usage screen.",
		o.ShowUsage,
		ExpectEnd,
	}
}
