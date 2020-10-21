package cli_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"fmt"
	"testing"
)

func TestConfigurationLoadDefault(t *testing.T) {
	var cfg cli.Configuration
	cfg.LoadDefault()
	errors.Assert(cfg.Concurrency == cli.DefaultConcurrency,
		fmt.Sprintf("Expected concurrency not found. Found:%d", cfg.Concurrency))
	errors.Assert(!cfg.Debug, "Expected debug to be false")
	errors.Assert(cfg.Delay == 0, "Expected delay not found.")
	errors.Assert(cfg.Depth == cli.DefaultDepth, "Expected depth not found.")
	errors.Assert(cfg.Dictionary == "", "Expected empty dictionary path/filename.")
	errors.Assert(cfg.Domain == "", "Expected empty domain string")
	errors.Assert(!cfg.Force, "Expected force to be false.")
	errors.Assert(cfg.Mode.IsNotSet(), "Expected notset mode")
	errors.Assert(cfg.Output == "", "Expected empty output filename.")
	errors.Assert(cfg.Pattern.String() == cli.DefaultFilterPattern, "Expected filter pattern not found.")
	errors.Assert(cfg.RecordTypes.String() == cli.DefaultDnsRecordTypes, "Expected dns record types not found.")
	errors.Assert(cfg.Timeout == cli.DefaultTimeout, "Expected timeout not found.")
}
