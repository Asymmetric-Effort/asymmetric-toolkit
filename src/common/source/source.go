package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/fifo"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
)

type Source struct {
	config Configuration
	isPaused       bool
	feed           fifo.Fifo
	counter        int
	logger 		   *logger.Logger
}
