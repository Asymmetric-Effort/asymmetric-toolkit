package source

import (
	"fmt"
)

func (o *Source) generateSequence() {
	/*
		Generate words as a sequence within a given keyspace (allowedChars).
		For example: A B C AA AB AC BA BB BC CA CB CC and so on where the
		length of the words is determined by config.WordSize.

		allowedRunes translates the allowedChars string into an array of bytes (runes)
		The position of the rune in this array is the value in the counter.
	 */
	allowedRunes :=
	runeCount:= uint8(len(allowedRunes))

	genString := func(counter []uint8) string {
		var runeString []rune = make([]rune, len(counter))
		for i := uint8(0); i < uint8(len(counter)); i++ {
			runeString[i] = allowedRunes[counter[i]]
		}
		s := string(runeString)
		return s
	}

	checkWordSize := func(wsz int) bool {
		return wsz <= int(o.config.WordSize)
	}
	checkWordCount := func() bool {
		return o.counter <= int(o.config.MaxWordCount)
	}
	checkWordLength := func(word *string) bool {
		return len(*word) != 0
	}
	for wordSize := 1; checkWordSize(wordSize) && checkWordCount(); wordSize++ {
		counter := make([]uint8, wordSize)
		//Generator factory function
		np := func() string {
			//then we create a counter variable which will roll through like an odometer.
			increment(&counter,0, runeCount)
			s:= genString(counter)
			fmt.Printf("wordSize:     %d\n", wordSize)
			fmt.Printf("allowedRunes: %v\n", allowedRunes)
			fmt.Printf("counter:      %v\n", counter)
			fmt.Printf("result:       '%s'\n", s)
			return s
		} //Execute generator factory function

		for word := np(); checkWordLength(&word) && checkWordCount(); word = np() {
			o.WaitIfPaused()
			o.feed <- word
			o.counter++
		}

	}
}
