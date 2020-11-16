package counter

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"math"
	"testing"
)

/*
	Increment 10(wordsize=1),100(wordsize=2),1000(wordsize=3),10000(wordsize=4)
	When we count through, we must account that the increasing wordsize will cause a leading character
	will be added with each carry-the-one operation.  This leading character will be the 'zero-character' in
    our sequence (the first character in the charset).  Accordingly we will observe the following progression--

		0,1,2,3,4,5,6,7,8,9

	followed by--

		00,01,02,03,04,05,06,07,08,09

	and so on.  Thus the number of words generated is

		n=f(s,b)

	where

		f(s,b) = b*(s-1) + f(s-1,b) if s > 0.
*/

func calcNumWords(s int, b int) (result int) {
	if s > 0 {
		result = int(math.Pow(float64(b), float64(s))) + calcNumWords(s-1, b)
	} else {
		result = 0
	}
	return
}

func TestCounter_Increment(t *testing.T) {
	const digits = "0123456789"
	tests := []struct {
		charset  string
		wordSize int
		expected int64
	}{
		{
			digits,
			1,
			10,
		}, {
			digits,
			2,
			110,
		}, {
			digits,
			3,
			1110,
		}, {
			digits,
			4,
			11110,
		},
	}
	func() {
		/*
			Test our calculations (pre-flight test)
		*/
		fmt.Println("Stage 1...")
		for _, test := range tests {
			n := int64(calcNumWords(test.wordSize, len(test.charset)))
			errors.Assert(n == test.expected, fmt.Sprintf("expected %d. Found %d", test.expected, n))
		}
		fmt.Println("Stage 1: OK")
	}()
	func() {
		/*
			Test the Counter.
		*/
		fmt.Println("Stage 2...")
		count := 0
		for _, test := range tests {
			var c Counter
			c.Setup(&test.charset, test.wordSize)
			for {
				count++
				//fmt.Printf("data:%v  count:%d\n", c.String(), count)
				if !c.Increment(0) {
					break
				}
			}
			expectedCount := int64(calcNumWords(test.wordSize, len(test.charset)))
			errors.Assert(int64(count) == expectedCount,
				fmt.Sprintf("Expected %d iterations.  Encountered:%d", count, expectedCount))
		}
		fmt.Println("Stage 2: OK")
	}()
}
