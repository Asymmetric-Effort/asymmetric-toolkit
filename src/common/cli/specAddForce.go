package cli

const (
	forceHelpText string = "Execute operations with force"
	forceDefault  string = ""
	forceArgLong         = "force"
)

func (o *Specification) AddForce() {
	//
	// Initialize the Argument object.
	//
	if o.Argument == nil {
		o.Argument = make(map[string]ArgumentDescriptor)
		o.Argument[""] = ArgumentDescriptor{}
	}
	//
	// We add a long argument for debug (--debug)
	//
	o.Argument[forceArgLong] = ArgumentDescriptor{
		flagForce,
		Boolean,
		forceDefault,
		forceHelpText,
		ParserFlag(forceArgLong),
		ExpectNone,
	}
}
