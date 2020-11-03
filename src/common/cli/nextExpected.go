package cli

/*
	The NextExpected value is used in a specification to instruct
	the CommandLine::Parse() function routing logic by declaring
	what type of token is to come after the current token.
*/
type NextExpected int

const (
	ExpectFlag  NextExpected = 0 // Expect a short (-c) or long (--char) flag.
	ExpectValue NextExpected = 1 // After a flag, we expect a matching value string.
	ExpectEnd   NextExpected = 2 // Expects the application to terminate.
	ExpectNone  NextExpected = 3 // A flag expects no value, which will run Parse() function then ExpectFlag
)
