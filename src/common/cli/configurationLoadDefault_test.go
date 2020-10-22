package cli_test

import (
	cli2 "asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestConfigurationLoadDefault(t *testing.T) {
	var cfg cli2.Configuration
	cfg.LoadDefault()
	errors.Assert(cfg.Concurrency == cli2.DefaultConcurrency,
		fmt.Sprintf("Expected concurrency not found. Found:%d", cfg.Concurrency))
	errors.Assert(!cfg.Debug, "Expected debug to be false")
	errors.Assert(cfg.Delay == 0, "Expected delay not found.")
	errors.Assert(cfg.Depth == cli2.DefaultDepth, "Expected depth not found.")
	errors.Assert(cfg.Dictionary == "", "Expected empty dictionary path/filename.")
	errors.Assert(cfg.Domain == "", "Expected empty domain string")
	errors.Assert(!cfg.Force, "Expected force to be false.")
	errors.Assert(cfg.Mode.IsNotSet(), "Expected notset mode")
	errors.Assert(cfg.Output == "", "Expected empty output filename.")
	errors.Assert(cfg.Pattern.String() == cli2.DefaultFilterPattern, "Expected filter pattern not found.")
	errors.Assert(cfg.RecordTypes.String() == cli2.DefaultDNSRecordTypes, "Expected dns record types not found.")
	errors.Assert(cfg.Timeout == cli2.DefaultTimeout, "Expected timeout not found.")
}
