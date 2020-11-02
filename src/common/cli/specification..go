package cli

/*
	The specification is the collection of arguments which are processed at runtime to determine how the
	commandline processor will interpret a given set of arguments.

	The specification consists of a set of metadata strings and a map of ArgumentDescriptors, each of which
	describes a specific argument and provides a parser function for this argument.  The CommandLine::Parse()
	function will use this ArgumentDescriptor::Parse() function to parse a specific argument's value when it
	is encountered and to sanitize/validate the value.
*/

type Specification struct {
	Author      string                        //Author name (used in usage notes and elsewhere)
	AuthorEmail string                        //Author email (used in usage notes and elsewhere)
	ProgramName string                        //program name (used in usage notes and elsewhere)
	Copyright   string                        //Copyright string (used in usage notes and elsewhere)
	Version     string                        //Application version (used in usage notes and elsewhere)
	Argument    map[string]ArgumentDescriptor //map of all arguments using the argument name as the key
}

func (o *Specification) ShowUsage() error {
	/*
		Calculate and show the usage message (all help messages).
	*/
	return nil
}
func (o *Specification) ShowVersion() error {
	/*
		Show the version string.
	*/
	return nil
}

func (o *Specification) AddUsage() {
	o.Argument["--help"] = ArgumentDescriptor{
		None,
		"",
		"Show help / usage screen.",
		o.ShowUsage,
	}
	o.Argument["-h"] = ArgumentDescriptor{
		None,
		"",
		"Show help / usage screen.",
		o.ShowUsage,
	}
}

func (o *Specification) AddVersion() {
	o.Argument["--version"] = ArgumentDescriptor{
		None,
		"",
		"Show version string",
		o.ShowVersion,
	}
	o.Argument["-v"] = ArgumentDescriptor{
		None,
		"",
		"Show version string",
		o.ShowVersion,
	}
}
