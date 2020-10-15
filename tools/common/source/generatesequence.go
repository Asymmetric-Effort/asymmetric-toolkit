package source

func (o *Source) generateSequence() {
	/*
		Generate words as a sequence within a given keyspace (allowedChars).
		For example: A B C AA AB AC BA BB BC CA CB CC and so on where the
		length of the words is determined by config.WordSize.
	 */

	for sz:=0; (sz<=o.config.WordSize.Get()) || (o.counter > int(o.config.MaxWordCount)); sz++{
		/*
			Iterate the size of our generated words from 0 to WordSize
			covering the entire keyspace with our allowedChars.

				number of results = allowedChars * WordSize
		 */
		o.WaitIfPaused()

		//Generator factory function
		np := func() func() string{

			//Setup internal state of the generator.
			allowedRunes := []rune(*o.allowedChars)
			wordSlice := make([]rune, sz)
			intArray := make([]int, len(wordSlice))

			//Returns generator function.
			return func() string {
				p := wordSlice[:len(intArray)]
				for i, xi := range intArray {
					p[i] = allowedRunes[xi]
				}
				for i := len(intArray) - 1; i >= 0; i-- {
					intArray[i]++
					if intArray[i] < len(allowedRunes) {
						break
					}
					intArray[i] = 0
					if i <= 0 {
						intArray = intArray[0:0]
						break
					}
				}
				o.counter++
				return string(p)
			}
		}() //Execute generator factory function

		for pwd := np(); len(pwd) != 0; pwd = np() {
			o.feed <- pwd
		}
	}
}
