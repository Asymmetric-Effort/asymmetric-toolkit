package misc

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestMiscIsIpAddr(t *testing.T) {
	count := 0
	for a := 0; a < 255; a += 5 {
		addr := fmt.Sprintf("%d.%d.%d.%d", a, a, a, a)
		errors.Assert(IsIpAddr(addr), fmt.Sprintf("Expected Ip (%s) to work", addr))
	}
	for a := 1; a < 255; a += 1 {
		for b := 0; b < 255; b += 1 {
			go func() {
				for c := 0; c < 255; c += 1 {
					addr := fmt.Sprintf("%d.%d.%d.254", a, b, c)
					errors.Assert(IsIpAddr(addr), fmt.Sprintf("Expected Ip (%s) to work", addr))
					count++
				}
			}()
		}
		fmt.Println("count:", count)
	}
	errors.Assert(!IsIpAddr("-1.-2.-10.-254"), "-1.-2.-10.-254 should be invalid")
	errors.Assert(!IsIpAddr("256.256.256.256"), "256.256.256.256 should be invalid")
	errors.Assert(!IsIpAddr("256.256.256"), "256.256.256 should be invalid")
	errors.Assert(!IsIpAddr("256.256.256:256"), "256.256.256:256 should be invalid")
}
