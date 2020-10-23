package cli

type Arguments map[string]interface{}

type Configuration struct {
	ProgramName string
	Terminate   bool
	args        Arguments
	spec        *Specification
}

const (
	expectFlag  = 0
	expectValue = 1
)
