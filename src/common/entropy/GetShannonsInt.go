package entropy

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"math"
)

func GetShannonsInt(data int) int {
	/*
	return func() int {

		totalLength := len(*input)
		var sum float64 = 0
		for _, i := range *input {
			for _, v := range *misc.CountNumberFrequency(i) {
				var f = float64(v) / float64(totalLength)
				sum += f * math.Log2(f)
			}
		}
		//return int(math.Ceil(sum*-1)) * int(totalLength)
	//}()

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
	freqMap:=misc.CountNumberFrequency(data)
	if data == 0 {
		return 0
	}
	for i := 0; i < 10; i++ {
		px := float64((*freqMap)[i]) / float64(misc.IntLen(data))
		if px > 0 {
			entropy += -px * math.Log2(px)
		}
	}
	return int(math.Ceil(entropy*100))
}
