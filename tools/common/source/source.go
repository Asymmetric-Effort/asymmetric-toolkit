package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
)

type Source struct {
	config         *cli.Configuration
	allowedChars   *string
	sourceBufferSz int
	isPaused       bool
	feed           chan<-string
	counter        int
}
