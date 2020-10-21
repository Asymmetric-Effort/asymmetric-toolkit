package errors

import (
	"fmt"
)

func Warn(s string) {
	/*
		Print a simple warning to the screen
	 */
	fmt.Printf("Warningf: %s\n", s)
}
