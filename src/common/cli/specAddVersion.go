package cli

func (o *Specification) AddVersion() {
	//ToDo: Do not overwrite if it exists.
	o.Argument[versionArgLong] = ArgumentDescriptor{
		flagVersion,
		None,
		"",
		"Show version string",
		o.ShowVersion,
		ExpectEnd,
	}
	o.Argument[versionArgShort] = ArgumentDescriptor{
		flagVersion,
		None,
		"",
		"Show version string",
		o.ShowVersion,
		ExpectEnd,
	}
}
