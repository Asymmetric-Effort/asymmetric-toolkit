package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/queues"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
)

type Source struct {
	Config         *cli.Configuration
	AllowedChars   *string
	SourceBufferSz int
	IsPaused       bool
	Feed           queues.Fifo
	Counter        int
	Logger         *logger.Logger
}
