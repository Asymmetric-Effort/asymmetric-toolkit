package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/deprecated_cli"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

const (
	TestSetupGeneratorTimeout time.Duration = 10 * time.Second
)

func TestSetupHappySequenceNullConfig(t *testing.T) {
	var s Source
	defer func() { recover() }()
	s.Setup(nil, deprecated_cli.SourceBufferSz, deprecated_cli.DnsChars)
}

func TestSetupHappySequenceBadBufferSize(t *testing.T) {
	var s Source
	var config deprecated_cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer","udp:127.0.0.1:53"}
	config.Parse(args)
	defer func() { recover() }()
	s.Setup(&config, 0, deprecated_cli.DnsChars)
}

func TestSetupHappySequenceBadChars(t *testing.T) {
	var s Source
	var config deprecated_cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer","udp:127.0.0.1:53"}
	config.Parse(args)
	defer func() { recover() }()
	s.Setup(&config, deprecated_cli.SourceBufferSz, "")
}

func TestSetupHappySequence(t *testing.T) {
	/*
		Setup
	*/
	testTimeout := func() {
		ticker := time.NewTicker(TestSetupGeneratorTimeout)
		defer func() { ticker.Stop() }()
		go func() { //Timeout
			fmt.Println("Starting timer to enforce timeout")
			for _ = range ticker.C {
				fmt.Println("Timeout Ex")
				t.Fatal("Timeout exceeded")
			}
		}()
	}
	for r := 1; r <= 4; r++ {
		fmt.Printf("wordsize: %d\n",r)
		var s Source
		var config deprecated_cli.Configuration
		args := []string{"--domain", "google.com", "--mode", "sequence", "--dnsServer","udp:127.0.0.1:53", "--wordSize", strconv.Itoa(r), "--maxWordCount", "100"}
		config.Parse(args)
		s.Setup(&config, deprecated_cli.SourceBufferSz, deprecated_cli.DnsChars)
		s.isPaused = false
		testTimeout()
		for s.HasData() {
			out := s.feed.Pop()
			t.Logf("Generator Output:%s", out)
		}
	}
}

func TestSetupHappyRandom(t *testing.T) {
	t.SkipNow()
	var s Source
	var config deprecated_cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "random", "--dnsServer","udp:127.0.0.1:53"}
	config.Parse(args)
	s.Setup(&config, deprecated_cli.SourceBufferSz, deprecated_cli.DnsChars)
	s.isPaused = false
	//go testTimeout(t,TestSetupGeneratorTimeout)
	for i := 0; i < 10; i++ {
		out := s.feed.Pop()
		if len(out) == 0 {
			t.Fatal("empty output")
		} else {
			t.Logf("Generator Output:%s", out)
		}
	}
}

func TestSetupHappyDictionary(t *testing.T) {
	t.SkipNow()
	var s Source
	var config deprecated_cli.Configuration
	//
	baseDir, err := os.Getwd()
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	//
	dictFile := filepath.Join(baseDir, "dictFile")
	errors.Assert(file.FileExists(dictFile), fmt.Sprintf("File not found:%s", dictFile))
	//
	args := []string{"--domain", "google.com", "--mode", "dictionary", "--dictionary", dictFile, "--dnsServer","udp:127.0.0.1:53"}
	config.Parse(args)
	s.Setup(&config, deprecated_cli.SourceBufferSz, deprecated_cli.DnsChars)
	s.isPaused = false
	// testTimeout(t,TestSetupGeneratorTimeout)
	for i := 0; i < 10; i++ {
		out := s.feed.Pop()
		if len(out) == 0 {
			t.Fatal("empty output")
		} else {
			t.Logf("Generator Output:%s", out)
		}
	}
}
