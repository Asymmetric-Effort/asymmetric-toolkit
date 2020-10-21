package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/random"
)

func (o *Source) GenerateRandom() {
	for o.Counter = 1; o.Counter <= int(o.Config.MaxWordCount); o.Counter++ {
		o.WaitIfPaused()
		o.Feed.Push(random.String(o.Config.WordSize.Get(),o.AllowedChars))
	}
}
