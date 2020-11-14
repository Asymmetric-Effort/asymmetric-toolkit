package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestLogWriterStdOut(t *testing.T) {
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)

	var log Logger
	log.Setup(&config)

	out := catchStdOut(t, func() {
		log.PrintEvent(&LogEventStruct{
			eventId: EventStd,
			time:    time.Now(),
			level:   Debug,
			tags:    nil,
			message: nil,
		})
	})
	msg := strings.Split(out, ":")
	word := strings.TrimSpace(msg[len(msg)-1])
	fmt.Println("out:", out)
	errors.Assert(word != "{}", fmt.Sprintf("Expected '{}'. Encountered: '%s'", word))
}
