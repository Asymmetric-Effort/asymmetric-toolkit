package cli

const (
	versionHelpText string = "Show version string"
	versionDefault  string = "false"
	versionArgLong  string = "version"
	versionArgShort string = "v"
)

func (o *Specification) AddVersion() {
	o.Initialize()
	//
	// We add a long argument for version (--version)
	//
	o.Argument[versionArgLong] = ArgumentDescriptor{
		FlagVersion,
		None,
		versionDefault,
		versionHelpText,
		o.ShowVersion,
		ExpectNone,
	}
}
