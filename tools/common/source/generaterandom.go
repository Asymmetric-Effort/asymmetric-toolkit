package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/random"
)

func (o *Source) generateRandom() {
	for o.counter = 1; o.counter <= int(o.config.MaxWordCount); o.counter++ {
		o.WaitIfPaused()
		o.feed.Push(random.String(o.config.WordSize.Get(),o.allowedChars))
	}
}
