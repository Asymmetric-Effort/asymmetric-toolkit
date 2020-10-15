package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/common/file"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	TestSetupGeneratorTimeout time.Duration = 10
)

func TestSetupHappySequenceNullConfig(t *testing.T) {
	var s Source
	defer func() { recover() }()
	s.Setup(nil, cli.SourceBufferSz, cli.DnsChars)
}

func TestSetupHappySequenceBadBufferSize(t *testing.T) {
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence"}
	config.Parse(args)
	defer func() { recover() }()
	s.Setup(&config, 0, cli.DnsChars)
}

func TestSetupHappySequenceBadChars(t *testing.T) {
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence"}
	config.Parse(args)
	defer func() { recover() }()
	s.Setup(&config, cli.SourceBufferSz, "")
}


func TestSetupHappySequence(t *testing.T) {
	t.SkipNow()
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "sequence"}
	config.Parse(args)
	s.Setup(&config, cli.SourceBufferSz, cli.DnsChars)
	s.isPaused = false
	go testTimeout(t,TestSetupGeneratorTimeout)
	for i := 0; i < 10; i++ {
		out := <-s.feed
		if len(out) == 0 {
			t.Fatal("empty output")
		} else {
			t.Logf("Generator Output:%s", out)
		}
	}
}

func TestSetupHappyRandom(t *testing.T) {
	t.SkipNow()
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "random"}
	config.Parse(args)
	s.Setup(&config, cli.SourceBufferSz, cli.DnsChars)
	s.isPaused = false
	go testTimeout(t,TestSetupGeneratorTimeout)
	for i := 0; i < 10; i++ {
		out := <-s.feed
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
	var config cli.Configuration
	//
	baseDir, err := os.Getwd()
	errors.Assert(err == nil, fmt.Sprintf("%v", err))
	//
	dictFile := filepath.Join(baseDir, "dictFile")
	errors.Assert(file.FileExists(dictFile), fmt.Sprintf("File not found:%s", dictFile))
	//
	args := []string{"--domain", "google.com", "--mode", "dictionary", "--dictionary", dictFile}
	config.Parse(args)
	s.Setup(&config, cli.SourceBufferSz, cli.DnsChars)
	s.isPaused = false
	go testTimeout(t,TestSetupGeneratorTimeout)
	for i := 0; i < 10; i++ {
		out := <-s.feed
		if len(out) == 0 {
			t.Fatal("empty output")
		} else {
			t.Logf("Generator Output:%s", out)
		}
	}
}
