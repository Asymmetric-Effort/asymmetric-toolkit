package source

func (o *Source) generateDictionary() {
	/*
		Read the o.dict file (observe pauseFlag)
	*/
	/*
	errors.Assert(o.config != nil, "Source::generateDictionary() encountered nil config")
	var dictionary DictionaryReader.Reader
	defer dictionary.Close()
	dictionary.Setup(dictionary.Read, o.config) //Configure the dictionary.

	dictionary.OpenRead()
	for line := scanner(); line != nil; line = scanner() {
		o.WaitIfPaused()
		o.queue.Push(*line)
		o.counter++
	}
	 */
}
