package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
)

type Source struct {
	config         *cli.Configuration
	allowedChars   *string
	sourceBufferSz int
	isPaused       bool
	feed           queues.Fifo
	counter        int
	logger         *logger.Logger
}
