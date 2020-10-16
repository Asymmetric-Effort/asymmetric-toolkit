package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"fmt"
	"testing"
	"time"
)

func TestSourceGenerateRandom(t *testing.T) {
	t.SkipNow()
	/*
		Setup
	*/
	const keyspace = "0123456789"
	var s Source
	var config cli.Configuration
	var seenBefore map[string]bool = make(map[string]bool, 1)
	args := []string{"--domain", "google.com", "--mode", "random", "--maxWordCount", "20"}
	config.Parse(args)
	s.config = &config
	s.config.WordSize = 2
	s.config.MaxWordCount = 1000
	s.allowedChars = func() *string { str := keyspace; return &str }()
	s.feed.Setup(cli.SourceBufferSz)
	go func() {
		<-time.After(time.Second * 120)
		t.Fatal("Terminating after timeout (120s)")
	}()
	/*
		Run Generator
	*/
	fmt.Println("Starting generator")
	go s.generateRandom()
	/*
		Analyze Result
	*/
	fmt.Println("Capture/analyze data from generator")
	for i := 1; i < 100; i++ {
		fmt.Printf("Processing results from generator %d", i)
		current := s.feed.Pop()
		if _, ok := seenBefore[current]; ok {
			t.Fatal("Collision detected")
		}
		seenBefore[current] = true
	}

}
