package cli

import (
	"fmt"
)

func (o *Specification) ShowVersion() (err error, val *Argument) {
	/*
		Show the version string.
	*/

	fmt.Printf("%s (%s)", o.ProgramName, o.Version)
	return nil, nil
}