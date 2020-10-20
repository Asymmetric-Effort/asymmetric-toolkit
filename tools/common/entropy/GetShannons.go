package entropy

import (
	"math"
	"strings"
)

func GetShannons(data string) int {
	/*
		Calculate the discrete Shannon entropy (H) of a given input string (n).

		Given a discrete random variable (x) that is the string of N "symbols" (total characters)
		consisting of n different characters (n=2 for binary), the shannon entropy of X in
		bits/symbols is:

					 _n_  count(i)       (  count(i) )
			H(X) = - \    ________ * log2(  ________ )
					 /__     N           (     N     )
					 i=1

		where count(i) is the count of characters n(i) (or frequency)
	*/
	var entropy float64 = 0
	if data == "" {
		return 0
	}
	for i := 0; i < 256; i++ {
		px := float64(strings.Count(data, string(byte(i)))) / float64(len(data))
		if px > 0 {
			entropy += -px * math.Log2(px)
		}
	}
	return int(math.Ceil(entropy*100))
}