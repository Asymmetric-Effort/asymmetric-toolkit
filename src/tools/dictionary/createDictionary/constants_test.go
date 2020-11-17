package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"testing"
)

func TestConstants(t *testing.T) {
	errors.Assert(DefaultLogLevel == logger.Info, "DefaultLogLevel mismatch")
}
