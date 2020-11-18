package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"testing"
)

func TestProcessSpecification(t *testing.T) {
	inputFilename:="testInput.txt"
	outputFilename:="testOutput.txt"
	args := []string{"--force", "--in", inputFilename, "--out", outputFilename}
	cfg, exit, err := ProcessSpecification(args)

	errors.Assert(cfg != nil, "Expect non nil config file.")
	errors.Assert(!exit, "Expect false exit flag.")
	errors.Assert(err == nil, "Expect no error.")

	errors.Assert(cfg.Log.Level == logger.Info, "Expected info level")
	errors.Assert(cfg.Log.Destination == logger.Stdout, "Expected stdout log destination")
	errors.Assert(cfg.Log.Settings == nil, "expect no settings.")

	errors.Assert(!cfg.Debug, "Expect false on debug")
	errors.Assert(cfg.Force, "Expect true on force")
	errors.Assert(cfg.InputFile==inputFilename, "Expected filename mismatch (input)")
	errors.Assert(cfg.InputFile==inputFilename, "Expected filename mismatch (input)")
}
