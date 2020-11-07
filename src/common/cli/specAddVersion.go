package cli

const (
	versionHelpText string = "Show version string"
	versionDefault  string = ""
	versionArgLong  string = "version"
	versionArgShort string = "v"
)

func (o *Specification) AddVersion() {
	//
	// Initialize the Argument object.
	//
	if o.Argument == nil {
		o.Argument = make(map[string]ArgumentDescriptor)
		o.Argument[""] = ArgumentDescriptor{}
	}
	//
	// We add a long argument for version (--version)
	//
	o.Argument[versionArgLong] = ArgumentDescriptor{
		flagVersion,
		None,
		versionDefault,
		versionHelpText,
		o.ShowVersion,
		ExpectEnd,
	}
	//
	// We add a short argument for version (-v)
	//
	o.Argument[versionArgShort] = ArgumentDescriptor{
		flagVersion,
		None,
		versionDefault,
		versionHelpText,
		o.ShowVersion,
		ExpectEnd,
	}
}
