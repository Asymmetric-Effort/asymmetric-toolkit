package errors

import "fmt"

func Assert(condition bool, msg string) bool {
	if condition {
		return true
	}
	panic(fmt.Sprintf("Assertion Error: %s", msg))
}
