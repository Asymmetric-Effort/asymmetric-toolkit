package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/fifo"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/deprecated_cli"
)

type Source struct {
	config         *deprecated_cli.Configuration
	allowedChars   *string
	sourceBufferSz int
	isPaused       bool
	feed           fifo.Fifo
	counter        int
	logger 		   *logger.Logger
}
