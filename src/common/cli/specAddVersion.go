package cli

const (
	versionHelpText string = "Show version string"
	versionDefault  string = ""
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
		ParserNoop(),
		ExpectNone,
	}
	//
	// We add a short argument for version (-v)
	//
	o.Argument[versionArgShort] = ArgumentDescriptor{
		FlagVersion,
		None,
		versionDefault,
		versionHelpText,
		ParserNoop(),
		ExpectNone,
	}
}
