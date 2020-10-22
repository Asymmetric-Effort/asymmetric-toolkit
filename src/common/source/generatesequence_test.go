package source_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	" asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"fmt"
	"testing"
)

func TestSourceGenerateSequence(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "0123456789"
	const maxWordSize = 9
	var s source.Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer", "udp:127.0.0.1:53"}
	config.Parse(args)
	s.config = &config

	for s.config.WordSize = 1; s.config.WordSize < maxWordSize; s.config.WordSize++ {
		s.config.MaxWordCount = 1000000000
		s.allowedChars = func() *string { str := keyspace; return &str }()
		s.feed.Setup(cli.SourceBufferSz * 1000)
		/*
			Run Generator
		*/
		fmt.Printf("Starting generator.  Size:%d of %d\n", s.config.WordSize, maxWordSize)
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
		fmt.Printf("Consumed queue of %d elements (last:%s)\n", expectedCount, last)
	}
	fmt.Println("TestSourceGenerateSequence...Done")
}

func TestSourceGenerateSequenceBinary(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "01"
	const maxWordSize = 9
	var s source.Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer","udp:127.0.0.1:53"}
	config.Parse(args)
	s.config = &config
	for s.config.WordSize = 1; s.config.WordSize < maxWordSize; s.config.WordSize++ {
		s.config.MaxWordCount = 1000000000
		s.allowedChars = func() *string { str := keyspace; return &str }()
		s.feed.Setup(cli.SourceBufferSz * 1000)
		/*
			Run Generator
		*/
		fmt.Printf("Starting generator.  Size:%d of %d\n", s.config.WordSize, maxWordSize)
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
		fmt.Printf("Consumed queue of %d elements (last:%s)\n", expectedCount, last)
	}
	fmt.Println("TestSourceGenerateSequenceBinary...Done")
}
