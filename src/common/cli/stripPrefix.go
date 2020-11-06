package cli

import (
	"regexp"
	"strings"
)

/*
	stripPrefix() is used by CommandLine::Parse() to strip the leading - or -- from a given flag
	and return the remaining string.  Thus, "-h" becomes "h" and "--help" becomes "help".

	Valid flag patterns include:
		-(char)
		--(>=2chars)
*/
func stripPrefix(arg *string) (isFlag bool, flag string) {
	//
	// Our regex should be satisfied by -<char> or --<string>
	//
	const regExFlagPattern = `^([-]{1}[a-zA-Z0-9]{1}|[-]{2}[a-zA-Z0-9]{2,24})$`
	re := regexp.MustCompile(regExFlagPattern)
	flag = strings.TrimLeft(*arg, "-")
	isFlag = re.MatchString(*arg) && (len(flag) > 0)
	return isFlag, flag
}
