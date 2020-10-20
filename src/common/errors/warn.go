package errors

import (
	"fmt"
)

func Warn(s string) {
	/*
		Print a simple warning to the screen
	 */
	fmt.Printf("Warning: %s\n", s)
}
