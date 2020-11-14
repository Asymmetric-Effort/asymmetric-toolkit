package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strings"
	"testing"
)

func TestLogWriterStdOut(t *testing.T) {
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)

	var log Logger
	log.Setup(&config)

	out := catchStdOut(t, func() {
		msg := []byte("test")
		log.logWriterStdOut(&msg)
	})
	msg := strings.Split(out, ":")
	word := strings.TrimSpace(msg[len(msg)-1])
	fmt.Printf("out:'%s'\n", out)
	errors.Assert(word == "test", fmt.Sprintf("Expected 'test'. Encountered: '%s'", word))
}
