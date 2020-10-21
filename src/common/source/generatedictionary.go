package source

/*
func (o *Source) generateDictionary() {
	//Read the o.dict file (observe pauseFlag)

	errors.Assert(o.config != nil,"Source::generateDictionary() encountered nil config")
	var dict reader.Reader
	defer dict.Close()
	scanner:=dict.Setup(o.config.Dictionary.OpenRead())
	for line:=scanner(); line != nil; line=scanner(){
		o.WaitIfPaused()
		o.feed.Push(*line)
		o.counter++
	}
}
*/
