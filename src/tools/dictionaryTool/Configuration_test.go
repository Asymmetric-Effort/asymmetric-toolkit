package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestProcessSpecification_Configuration(t *testing.T) {
	var config Configuration
	errors.Assert(config.dictionary == "", "expected empty string")
}