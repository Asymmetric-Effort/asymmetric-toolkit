package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/LogDestination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestLoggerWarningError(t *testing.T) {
	var log Logger
	var config cli.CommandLine
	config.Log.Destination.Set(LogDestination.Stdout)
	config.Log.Level.Set(level.Warning)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Warning("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(WARN\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerWarningCritical(t *testing.T) {
	var log Logger
	var config cli.CommandLine
	config.Log.Destination.Set(LogDestination.Stdout)
	config.Log.Level.Set(level.Warning)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Critical("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(CRIT\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}

func TestLoggerWarningDebug(t *testing.T) {
	var log Logger
	var config cli.CommandLine
	config.Log.Destination.Set(LogDestination.Stdout)
	config.Log.Level.Set(level.Debug)
	out := catchStdOut(t, func() {
		log.Setup(&config)
		log.Warning("Test")
	})
	pattern := `^\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+ [-+]{1}[0-9]+ [A-Z]+ m=[+-][0-9]+\.[0-9]+\]\[Logger\]\(WARN\): Test$`
	re := regexp.MustCompile(pattern)
	errors.Assert(re.MatchString(strings.TrimRight(out, "\n")), fmt.Sprintf("Pattern failed to match: %s", out))
}
