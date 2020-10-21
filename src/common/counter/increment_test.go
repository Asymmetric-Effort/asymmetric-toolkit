package counter

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"math"
	"testing"
)

const (
	charset="0123456789"
)
func TestCounter_IncrementHappy(t *testing.T) {
	for wordsize:=1;wordsize<4;wordsize++ {
		var c Counter
		count:=1
		for c.Setup(charset, wordsize) ; !c.Increment(0); count++ {
			fmt.Printf("data:%v\n", *c.data)
		}
		expectedCount := int64(math.Pow(float64(len(charset)), float64(wordsize)))
		errors.Assert(int64(count) == expectedCount, fmt.Sprintf("Expected %d iterations.  Encountered:%d", count, expectedCount))

	}
}
