package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"fmt"
	"math"
	"strconv"
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
	s.config.WordSize = 2
	s.config.MaxWordCount = 1000
	s.allowedChars = func() *string { str := keyspace; return &str }()
	s.feed.Setup(cli.SourceBufferSz*10)
	/*
		Run Generator
	*/
	fmt.Println("Starting generator")
	s.generateSequence()
	fmt.Printf("Finished generating.  Number items in channel: %d\n", s.feed.Length())
	/*
		Analyze Result
	*/
	fmt.Println("Capture/analyze data from generator")
	count := 0
	expectedCount := int(math.Pow(float64(len(keyspace)), float64(s.config.WordSize)))
	for i:=0;i<expectedCount;i++ {
		fmt.Printf("Processing...  ChanSz: %d Count:%d\n", s.feed.Length(), count)
		if s.HasData() {
			element := s.feed.Pop()
			_, err := strconv.Atoi(element)
			if err != nil {
				panic(err)
			}
			count++
		}
	}
	errors.Assert(count == expectedCount, fmt.Sprintf("Count (%d) does not match expectedCount (%d) does not match.", count, expectedCount))
}
