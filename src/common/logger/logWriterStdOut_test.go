package logger_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"fmt"
	"strings"
	"testing"
)

func TestLogWriterStdOut(t *testing.T) {
	var log logger.Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(level.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Printf(level.Debug, "%s", "Test")
	})
	msg := strings.Split(out, ":")
	word := strings.TrimSpace(msg[len(msg)-1])
	errors.Assert(word == "Test", fmt.Sprintf("Expected 'Test'. Encountered: '%s'", word))
}
