package errors

import (
	"fmt"
)

func Warn(s string) {
	fmt.Printf("Warning: %s\n", s)
}
