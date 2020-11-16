package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestProcessSpecification_Configuration(t *testing.T) {
	var config Configuration
	errors.Assert(config.Concurrency == 0, "expected 0")
	errors.Assert(config.Delay == 0, "expected 0")
	errors.Assert(config.Depth == 0, "expected 0")
	errors.Assert(config.DnsServer == "", "expected empty string")
	errors.Assert(config.DnsRecordTypes == "", "expected empty string")
	errors.Assert(config.Domain == "", "expected empty string")
	errors.Assert(config.Output == "", "expected empty string")
	errors.Assert(config.Timeout == 0, "expected 0")
}
