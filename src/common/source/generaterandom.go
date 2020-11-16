package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/random"
)

func (o *Source) generateRandom() {
	for o.counter = 1; o.counter <= int(o.config.MaxWordCount); o.counter++ {
		o.WaitIfPaused()
		o.queue.Push(random.String(o.wordSize, &o.config.AllowedChars))
	}
}
