package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/counter"
)

func (o *Source) GenerateSequence() {
	/*
		Generate words as a sequence within a given keyspace (allowedChars).
		For example: A B C AA AB AC BA BB BC CA CB CC and so on where the
		length of the words is determined by config.WordSize.
	 */
	checkWordCount := func() bool {
		return o.Counter <= int(o.Config.MaxWordCount)
	}
	for wordSize := 1; (wordSize <= int(o.Config.WordSize)) && checkWordCount(); wordSize++ {
		var generator counter.Counter
		generator.Setup(*o.AllowedChars, wordSize)
		//Iterate through the field of wordSize chars.
		o.Feed.Push(generator.String())
		for generator.Increment(0) && checkWordCount() {
			//Produce a new word for each one and push to the feed.
			o.Feed.Push(generator.String())
			o.Counter++
		}
	}
}
