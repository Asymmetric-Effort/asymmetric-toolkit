package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
)

func (o *Source) Setup(config *cli.Configuration, bufferSz int, allowedChars string) {
	/*
	errors.Assert(config != nil, "Encountered nil configuration in Source::Setup()")
	errors.Assert(bufferSz > 1, "Expected sourceBufferSz > 1")
	errors.Assert(allowedChars != "", "Expected non-empty string in allowedChars")

	o.config = config
	o.feed.Setup(bufferSz)
	o.allowedChars = &allowedChars
	o.isPaused = true //By default the feed will be paused until the owner unpauses it.

	switch {
	case config.Mode.IsSequence():
		go o.generateSequence()
	case config.Mode.IsRandom():
		go o.generateRandom()
	case config.Mode.IsDictionary():
		go o.generateDictionary()
	default:
		panic("Source::Setup() encountered Mode NotSet")
	}

	 */
}
