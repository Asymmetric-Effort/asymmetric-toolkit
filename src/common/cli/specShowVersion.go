package cli
/*

 */

import (
	"fmt"
)

func (o *Specification) ShowVersion(arg *string) (err error, val *Argument) {
	/*
		Show the version string.
	*/
	fmt.Printf("%s (%s)\n", o.ProgramName, o.Version)
	return nil, nil
}