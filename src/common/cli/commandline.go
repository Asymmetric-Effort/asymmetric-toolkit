package cli

type Arguments map[string] interface{}

type Configuration struct {
	Terminate bool
	args Arguments
	spec *Specification
}

const (
	expectFlag  = 0
	expectValue = 1
)
