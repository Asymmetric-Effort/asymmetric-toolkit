package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/fifo"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
)

type Source struct {
	config         *cli.Configuration
	allowedChars   *string
	sourceBufferSz int
	isPaused       bool
	feed           fifo.Fifo
	counter        int
	logger 		   *logger.Logger
}
