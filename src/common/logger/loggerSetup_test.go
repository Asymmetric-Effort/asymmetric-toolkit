package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogger_Setup(t *testing.T) {
	var log Logger
	var config = Configuration{
		Level:       Debug,
		Settings:    nil,
		Destination: Stdout,
	}
	log.Setup(&config)
	errors.Assert(log.Level == Debug, "Expect Debug")
	errors.Assert(log.Destination == Stdout, "expect Stdout")
	errors.Assert(log.Settings == nil, "expect nil")
	errors.Assert(log.Writer != nil, "expect nil")
	errors.Assert(&log.tags != nil, "expect non-nil address")
}
