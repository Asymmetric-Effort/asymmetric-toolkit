package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"fmt"
	"testing"
)

func TestSourceGenerateSequence(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "0123456789"
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence"}
	config.Parse(args)
	s.config = &config
	for s.config.WordSize = 1; s.config.WordSize < 9; s.config.WordSize++ {

		s.config.MaxWordCount = 1000000000
		s.allowedChars = func() *string { str := keyspace; return &str }()
		s.feed.Setup(cli.SourceBufferSz * 1000)
		/*
			Run Generator
		*/
		fmt.Printf("Starting generator.  Size:%d\n", s.config.WordSize)
		s.generateSequence()
		s.feed.Close()
		fmt.Printf("Finished generating.  Number items in channel: %d\n", s.feed.Length())
		/*
			Analyze Result
		*/
		expectedCount := s.feed.Length()
		last := ""
		for i := 0; i <= expectedCount; i++ {
			if s.feed.Length() > 0 {
				last = s.feed.Pop()
			}
		}
		errors.Assert(s.feed.Length() == 0, "Expected to have consumed all elements")
		fmt.Printf("Consumed queue of %d elements (last:%s)", expectedCount, last)
	}
}
