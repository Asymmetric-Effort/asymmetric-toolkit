package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestLoggerErrorError(t *testing.T) {
	var log Logger
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Error)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Error("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(ERROR\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerErrorCritical(t *testing.T) {
	var log Logger
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Error)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Critical("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(CRIT\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerErrorDebug(t *testing.T) {
	var log Logger
	var config Configuration
	config.Destination.Set(Stdout)
	config.Level.Set(Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Error("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(ERROR\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))

}
