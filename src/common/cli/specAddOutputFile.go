package cli
/*
	Specification::AddOutputFile() implements the --in <filename>.
*/
const (
	outFileHelpText string = "Specify an output file"
	outFileDefault  string = ""

	outFileArgLong  = "out"
)

func (o *Specification) AddOutputFile() {
	o.Initialize()
	//
	// We add a long argument for the output file (--out)
	//
	o.Argument[outFileArgLong] = ArgumentDescriptor{
		FlagOutputFile,
		String,
		outFileDefault,
		outFileHelpText,
		ParserString(),
		ExpectValue,
	}
}
