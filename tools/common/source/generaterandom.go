package source

import (
	"math/rand"
	"time"
)

func (o *Source) generateRandom() {
	keySpace := []rune(*o.allowedChars)
	numRunes := len(*o.allowedChars)
	word := make([]rune, o.config.WordSize.Get())

	for o.counter = 1; o.counter <= int(o.config.MaxWordCount); o.counter++ {
		o.WaitIfPaused()
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano() + int64(o.counter)))
		for i := range word {
			word[i] = keySpace[seededRand.Intn(numRunes)]
		}
		o.feed.Push(string(word)) //ToDo: Add FiFO queue with blocking send
	}
}
