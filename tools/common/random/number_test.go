package random

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestNumberLarge(t *testing.T) {
	rounds := 1000000
	runs := 10
	low := 1
	high := 4294967295
	for run := runs; run > 0; run-- {
		fmt.Println("Run", run, "/", runs)
		above := 0
		frequency := make(map[int]int)
		for i := 0; i < rounds; i++ {
			n := Number(low, high)
			if entropy.HighEntropyInt(n) {
				above++
			}
			if _, ok := frequency[n]; ok {
				frequency[n]++
			} else {
				frequency[n] = 1
			}
		}

		pct := 100 * float64(above) / float64(rounds)
		errors.Assert(pct > 99.9, fmt.Sprintf("Expected 99.9%% random numbers had high entropy. "+
			"Encountered: %f", pct))

		above = 0
		for _, v := range frequency {
			if v > 1 {
				above++
			}
		}
		pct = 100 * float64(above) / float64(rounds)
		errors.Assert(pct < 0.1, fmt.Sprintf("Expected 0.02%% with frequency over 1."+
			"Encountered: %f", pct))
	}
}
