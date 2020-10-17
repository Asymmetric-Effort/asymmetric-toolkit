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
	for length := 20; length <= 100; length++ {
		for count := 1; count < 10000; count++ {
			data := String(100, &keySpace)
			score := entropy.GetShannons(data)
			errors.Assert(score > entropy.HighEntropyThreshold, "Expected entropy over threshold.")
			fmt.Printf(".")
		}
	}
	fmt.Println("done")
}
