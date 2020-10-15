package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/common/file/dictionary"
)

func (o *Source) generateDictionary() {
	/*
		Read the o.dict file (observe pauseFlag)
	*/
	errors.Assert(o.config != nil,"Source::generateDictionary() encountered nil config")
	var dict dictionary.Reader
	defer dict.Close()
	scanner:=dict.Setup(o.config.Dictionary.OpenRead())
	for line:=scanner(); line != nil; line=scanner(){
		o.WaitIfPaused()
		o.feed <- *line
		o.counter++
	}
}
