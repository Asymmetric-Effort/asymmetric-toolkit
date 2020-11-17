package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSourceGenerateSequence(t *testing.T) {
	/*
		Setup
	*/
	const keyspace10 = "0123456789"
	const keyspace2 = "01"
	const bufferSize = 1048576

	tests := []struct {
		keyspace string
	}{
		{
			keyspace2,
		}, {
			keyspace10,
		},
	}

	for i, test := range tests {
		var s Source
		var config Configuration
		config.MaxWordCount = 1000000000
		config.MinWordSize = 1
		config.MaxWordSize = 5
		config.BufferSize = bufferSize
		fmt.Printf("Test #%d: START\n",i)
		s.Setup(&config)
		for s.wordSize = s.config.MinWordSize; s.wordSize < s.config.MaxWordSize; s.wordSize++ {

			s.config.AllowedChars = test.keyspace

			s.queue.Setup(s.config.BufferSize)
			/*
				Run Generator
			*/
			fmt.Printf("Starting Generator.  Size:%d of %d\n", s.wordSize, s.config.MaxWordSize)
			s.generateSequence()
			s.queue.Close()
			fmt.Printf("Finished generating.  Number items in channel: %d\n", s.queue.Length())
			/*
				Analyze Result
			*/
			expectedCount := s.queue.Length()
			last := ""
			for i := 0; i <= expectedCount; i++ {
				if s.queue.Length() > 0 {
					last = s.queue.Pop()
				}
			}
			errors.Assert(s.queue.Length() == 0, "Expected to have consumed all elements")
			fmt.Printf("Consumed queue of %d elements (last:%s)\n", expectedCount, last)
		}
		fmt.Printf("Test #%d: OK\n",i)
	}
	fmt.Println("TestSourceGenerateSequence...Done")
}
