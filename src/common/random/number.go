package random

import (
	"math/rand"
	"time"
)

func Number(low int, high int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(high-low+1) + low
}
