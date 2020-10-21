package errors

import (
	"fmt"
	"os"
)

func Fatal(c int, s string) {
	fmt.Printf("Fatal errors: %s\n", s)
	os.Exit(c)
}
