package source

func (o *Source) generateRandom() {
	o.WaitIfPaused()
	for sz:=0;(sz<=o.config.WordSize.Get()) || (o.counter> int(o.config.MaxWordCount)); sz++{
		o.WaitIfPaused()
	}
}
