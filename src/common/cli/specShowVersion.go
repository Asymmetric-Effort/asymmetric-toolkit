package cli
/*

 */

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
)

func (o *Specification) ShowVersion(arg *string) (err error, val *Argument) {
	/*
		Show the version string.
	*/
	errors.Assert(arg != nil, "Expected non nil argument.")
	fmt.Printf("%s (%s)\n", o.ProgramName, o.Version)
	return nil, nil
}