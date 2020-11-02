package cli

import "strings"

/*
	stripPrefix() is used by CommandLine::Parse() to strip the leading - or -- from a given flag
	and return the remaining string.  Thus, "-h" becomes "h" and "--help" becomes "help".

	Valid flag patterns include:

		-(char)
		--(>=2chars)
*/
func stripPrefix(arg *string) (isFlag bool, flag string) {
	isFlag = ((((*arg)[0:1] == "-") && (len(*arg) == 2)) || ((*arg)[0:2] == "--") && (len(*arg) > 3)) &&
		((*arg)[0:3] != "---")
	flag = strings.TrimLeft(*arg, "-")
	return
}
