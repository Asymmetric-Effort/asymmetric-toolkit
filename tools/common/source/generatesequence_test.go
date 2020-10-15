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
	t.SkipNow()
	const keyspace = "0123456789"
	var s Source
	var config cli.Configuration
	/*
		Setup
	*/
	args := []string{"--domain", "google.com", "--mode", "sequence"}
	config.Parse(args)
	s.config = &config
	s.config.WordSize = 2
	s.config.MaxWordCount = 1000
	s.allowedChars = func() *string { str := keyspace; return &str }()
	s.feed = make(chan string, cli.SourceBufferSz)
	go testTimeout(t, TestSetupGeneratorTimeout)
	fmt.Println("Starting generator")
	/*
		Run Generator
	*/
	go s.generateSequence()
	/*
		Analyze Result
	*/
	fmt.Println("Capture data from generator")
	count := 0
	for element := <-s.feed; len(s.feed) != 0; element = <-s.feed {
		count++
		current, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		errors.Assert(current <= count, "Count should be greater than current.")
	}
	expectedCount:=int(math.Pow(float64(len(keyspace)),float64(s.config.WordSize)))
	errors.Assert(count == expectedCount, fmt.Sprintf("Count (%d) does not match expectedCount (%d) does not match.",count,expectedCount))
}
