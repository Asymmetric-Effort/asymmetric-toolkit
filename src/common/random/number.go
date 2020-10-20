package random

import (
	"math/rand"
	"time"
)

func Number(low int, high int) int {
	/*
		return a random integer between low and high
	 */
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(high-low+1) + low
}
