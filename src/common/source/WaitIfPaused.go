package source

import "time"

const PauseCheckDelay time.Duration = 5

func (o *Source) WaitIfPaused() {
	for o.IsPaused {
		<-time.After(time.Second * PauseCheckDelay)
	}
}
