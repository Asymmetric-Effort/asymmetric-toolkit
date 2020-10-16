package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"fmt"
	"testing"
)

func TestSourceGenerateRandom(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "0123456789"
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "random", "--dnsServer", "udp:127.0.0.1:53", "--maxWordCount", "20"}
	config.Parse(args)
	s.config = &config
	s.config.WordSize = 4
	s.config.MaxWordCount = 1000
	s.allowedChars = func() *string { str := keyspace; return &str }()
	s.feed.Setup(cli.SourceBufferSz)
	/*
		Run Generator
	*/
	fmt.Println("Starting generator")
	s.generateRandom()
	s.feed.Close()
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
