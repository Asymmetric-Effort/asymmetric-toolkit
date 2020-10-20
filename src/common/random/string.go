package random

import (
	"math/rand"
	"os"
	"time"
)

func String(sz int,chars *string) string {
	/*
		Return a random string of sz length using the chars character set.
	 */
	keySpace := []rune(*chars)
	word := make([]rune, sz)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano() + int64(os.Getpid())))
	for i := range word {
		word[i] = keySpace[seededRand.Intn(len(keySpace))]
	}
	return string(word)
}
