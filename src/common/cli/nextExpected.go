package cli
/*
	The NextExpected value is used in a specification to instruct
	the CommandLine::Parse() function routing logic by declaring
	what type of token is to come after the current token.
 */
type NextExpected int

const (
	ExpectFlag  NextExpected = 0
	ExpectValue NextExpected = 1
	ExpectEnd   NextExpected = 2
	ExpectNone  NextExpected = 3
)