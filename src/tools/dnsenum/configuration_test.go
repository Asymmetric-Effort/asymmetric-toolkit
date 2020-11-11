package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestProcessSpecification_Configuration(t *testing.T) {
	var config Configuration
	errors.Assert(config.Concurrency == 0, "expected 0")
	errors.Assert(config.delay == 0, "expected 0")
	errors.Assert(config.depth == 0, "expected 0")
	errors.Assert(config.dictionary == "", "expected empty string")
	errors.Assert(config.dnsServer == "", "expected empty string")
	errors.Assert(config.dnsRecordTypes == "", "expected empty string")
	errors.Assert(config.domain == "", "expected empty string")
	errors.Assert(config.maxWordCount == 0, "expected 0")
	errors.Assert(config.minWordSize == 0, "expected 0")
	errors.Assert(config.maxWordSize == 0, "expected 0")
	errors.Assert(config.output == "", "expected empty string")
	errors.Assert(config.pattern == "", "expected empty string")
	errors.Assert(config.source == "", "expected empty string")
	errors.Assert(config.timeout == 0, "expected 0")
}