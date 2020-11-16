package source

import "asymmetric-effort/asymmetric-toolkit/src/common/counter"

func (o *Source) generateSequence() {
	/*
		Generate words as a sequence within a given keyspace (allowedChars).
		For example: A B C AA AB AC BA BB BC CA CB CC and so on where the
		length of the words is determined by config.WordSize.
	*/
	checkWordCount := func() bool {
		return o.counter <= int(o.config.MaxWordCount)
	}
	for wordSize := 1; (wordSize <= int(o.wordSize)) && checkWordCount(); wordSize++ {
		var generator counter.Counter
			generator.Setup(&o.config.AllowedChars, wordSize)
		//Iterate through the field of wordSize chars.
		o.queue.Push(generator.String())
		for generator.Increment(0) && checkWordCount() {
			//Produce a new word for each one and push to the feed.
			o.queue.Push(generator.String())
			o.counter++
		}
	}
}
