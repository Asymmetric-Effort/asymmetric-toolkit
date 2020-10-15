package counter

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"math"
	"testing"
)

const (
	charset="0123456789"
)
func TestCounter_IncrementHappy10(t *testing.T) {
	var c Counter
	const wordsize = 1
	c.Setup(charset,wordsize)
	count:=1
	for {
		fmt.Printf("data:%v\n",*c.data)
		if !c.Increment(0) {
			break
		}
		count++
	}
	expectedCount:=int64(math.Pow(float64(len(charset)), float64(wordsize)))
	errors.Assert(int64(count) == expectedCount, fmt.Sprintf("Expected %d iterations.  Encountered:%d",count,expectedCount))
}
func TestCounter_IncrementHappy100(t *testing.T) {
	var c Counter
	const wordsize = 2
	c.Setup(charset,wordsize)
	count:=1
	for {
		fmt.Printf("data:%v\n",*c.data)
		if !c.Increment(0) {
			break
		}
		count++
	}
	expectedCount:=int64(math.Pow(float64(len(charset)), float64(wordsize)))
	errors.Assert(int64(count) == expectedCount, fmt.Sprintf("Expected %d iterations.  Encountered:%d",count,expectedCount))
}
func TestCounter_IncrementHappy1000(t *testing.T) {
	var c Counter
	const wordsize = 3
	c.Setup(charset,wordsize)
	count:=1
	for {
		fmt.Printf("data:%v\n",*c.data)
		if !c.Increment(0) {
			break
		}
		count++
	}
	expectedCount:=int64(math.Pow(float64(len(charset)), float64(wordsize)))
	errors.Assert(int64(count) == expectedCount, fmt.Sprintf("Expected %d iterations.  Encountered:%d",count,expectedCount))
}