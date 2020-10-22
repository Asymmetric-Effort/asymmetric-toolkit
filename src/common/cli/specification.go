package cli

/*
	The command line specification allows any tool to define its commandline
	expectations as a map of terms, where each term is a single '--flag value'
	pair.
*/
type Specification map[string]Term

type Term struct {
	Required    bool //Argument is required
	ExpectValue bool // Does flag require a following value (e.g. "--flag myValue")
	Validation  ValidationFunc
	Help        string
}
