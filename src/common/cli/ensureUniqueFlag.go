package cli

import "fmt"

/*
	Specification::EnsureUniqueFlagId() is a method which scans the arguments in a given Specification
	to ensure all Argument FlagId properties are unique.  The application would most likely break if there
	were a duplicate FlagIds.
*/

func (o *Specification) EnsureUniqueFlagId() (err error) {

	seenBefore := make(map[ArgumentFlag]bool)

	for name, arg := range o.Argument {

		ok := seenBefore[arg.FlagId]

		if ok {
			return fmt.Errorf("duplicate flagId (%d) in Specification (%s)", arg.FlagId, name)
		}

		seenBefore[arg.FlagId] = true

	}

	return nil

}
