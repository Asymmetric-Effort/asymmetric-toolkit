package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/fifo"
	"asymmetric-effort/asymmetric-toolkit/tools/common/logger"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
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
