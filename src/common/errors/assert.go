package errors

import "fmt"

func Assert(condition bool, msg string) bool {
	/*
		This is an assertion which will test a condition
		and if that condition is not true, the application
		will panic and display the given message.
	 */
	if condition {
		return true
	}
	panic(fmt.Sprintf("Assertion Errorf: %s", msg))
}
