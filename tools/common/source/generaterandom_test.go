package source

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/entropy"
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/cli"
	"fmt"
	"testing"
)

func TestSourceGenerateRandom(t *testing.T) {
	/*
		Setup
	*/
	const keyspace = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var s Source
	var config cli.Configuration
	args := []string{"--domain", "google.com", "--mode", "random", "--dnsServer", "udp:127.0.0.1:53", "--maxWordCount", "20"}
	config.Parse(args)
	s.config = &config
	s.config.WordSize = 100
	s.config.MaxWordCount = 100
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
	word := ""
	totalShannons := 0
	minShannons := entropy.HighEntropyThreshold
	maxShannons := 0
	count := 0
	for i := 0; i <= expectedCount; i++ {
		if s.feed.Length() > 0 {
			word = s.feed.Pop()
			count++
			shannons := entropy.GetShannons(word)
			totalShannons += shannons
			if shannons < minShannons {
				minShannons = shannons
			}
			if shannons > maxShannons {
				maxShannons = shannons
			}
			fmt.Println("Shannons:", totalShannons, "shannons.  Avg", totalShannons/count)
		}
	}
	avgShannons := totalShannons / count
	errors.Assert(s.feed.Length() == 0, "Expected to have consumed all elements")
	errors.Assert(float64(minShannons) > (0.90 * float64(entropy.HighEntropyThreshold)), fmt.Sprintf("Expected high min entropy (%d)",minShannons))
	errors.Assert(avgShannons > entropy.HighEntropyThreshold, "Expected high average entropy")
	errors.Assert(maxShannons > entropy.HighEntropyThreshold, "Expected high max entropy")
	errors.Assert(count == expectedCount, "Expected count match")
	fmt.Printf("Consumed queue of %d elements", expectedCount)
}
