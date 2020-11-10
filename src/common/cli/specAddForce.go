package cli

const (
	forceHelpText string = "Execute operations with force"
	forceDefault  string = ""
	forceArgLong         = "force"
)

func (o *Specification) AddForce() {
	o.Initialize()
	//
	// We add a long argument for debug (--debug)
	//
	o.Argument[forceArgLong] = ArgumentDescriptor{
		FlagForce,
		Boolean,
		forceDefault,
		forceHelpText,
		ParserFlag(forceArgLong),
		ExpectNone,
	}
}
