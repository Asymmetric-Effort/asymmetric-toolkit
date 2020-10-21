package logger_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestLoggerCriticalCritical(t *testing.T) {
	var log Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(logLevel.Critical)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Critical("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(CRIT\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerCriticalDebug(t *testing.T) {
	var log Logger
	var config cli.Configuration
	config.Log.Destination.Set(destination.Stdout)
	config.Log.Level.Set(logLevel.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Critical("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(CRIT\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}
