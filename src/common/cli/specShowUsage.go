package cli
/*
	ShowUsage is a ParserFunction, following the same pattern.  However, this parser function will not actually
	parse any argument or create any value.  It simply calculates then displays program usage information (flags, etc).
 */

import (
	"fmt"
)

func (o *Specification) ShowUsage(arg *string) (err error, val *Argument) {
	/*
		Calculate and show the usage message (all help messages).
	*/
	//ToDo: calculate usage from specification.
	fmt.Println("Not implemented.")
	return nil, nil
}
