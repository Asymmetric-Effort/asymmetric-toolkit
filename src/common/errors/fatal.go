package errors

import (
	"fmt"
	"os"
)

func Fatal(c int, s string) {
	/*
		Show the given message and then exit with the given integer exit code.
	 */
	fmt.Printf("Fatal errors: %s\n", s)
	os.Exit(c)
}
