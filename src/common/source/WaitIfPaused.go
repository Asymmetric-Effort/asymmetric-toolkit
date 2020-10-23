package source

import "time"

const pauseCheckDelay time.Duration = 5

func (o *Source) WaitIfPaused() {
	for o.isPaused {
		<-time.After(time.Second * pauseCheckDelay)
	}
}
