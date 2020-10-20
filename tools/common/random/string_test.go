package random

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestRandomStringLength(t *testing.T) {
	keySpace := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	func() {
		s := String(100, &keySpace)
		errors.Assert(len(s) == 100, "Expected 100 characters")
	}()
	for length := 8; length <= 100; length++ {
		for count := 1; count < 10000; count++ {
			data := String(length, &keySpace)
			errors.Assert(entropy.HighEntropy(data), "Expected entropy over threshold.")
		}
		fmt.Printf(".")
	}
	fmt.Println("done")
}
