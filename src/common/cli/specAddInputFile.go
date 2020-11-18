package cli
/*
	Specification::AddInputFile() implements the --in <filename>.
*/
const (
	inFileHelpText string = "Specify an input file"
	inFileDefault  string = ""

	inFileArgLong  = "in"
)

func (o *Specification) AddInputFile() {
	o.Initialize()
	//
	// We add a long argument for the input file. (--in)
	//
	o.Argument[inFileArgLong] = ArgumentDescriptor{
		FlagInputFile,
		String,
		inFileDefault,
		inFileHelpText,
		ParserString(),
		ExpectValue,
	}
}
