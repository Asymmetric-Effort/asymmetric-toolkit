package cli

import (
	args "asymmetric-effort/asymmetric-toolkit/src/common/cli/args"
)

type CommandLine struct {
	ProgramName string
	Terminate   bool
	args        args.Arguments
	spec        *Specification
}
