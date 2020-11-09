package cli

/*
	Specification::Initialization()
*/
func (o *Specification) Initialize() {
	//
	// Initialize the Argument object.
	//
	if o.Argument == nil {
		o.Argument = make(map[string]ArgumentDescriptor)
		o.Argument[""] = ArgumentDescriptor{}
	}
}
