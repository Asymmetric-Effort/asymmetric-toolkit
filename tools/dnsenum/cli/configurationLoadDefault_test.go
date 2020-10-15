package cli

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"testing"
)

func TestConfigurationLoadDefault(t *testing.T) {
	var cfg Configuration
	cfg.LoadDefault()
	errors.Assert(cfg.Concurrency == defaultConcurrency, fmt.Sprintf("Expected concurrency not found. Found:%d", cfg.Concurrency))
	errors.Assert(!cfg.Debug, "Expected debug to be false")
	errors.Assert(cfg.Delay == 0, "Expected delay not found.")
	errors.Assert(cfg.Depth == defaultDepth, "Expected depth not found.")
	errors.Assert(cfg.Dictionary == "", "Expected empty dictionary path/filename.")
	errors.Assert(cfg.TargetServers.String() == "", "Expected dns servers not found.")
	errors.Assert(cfg.Domain == "", "Expected empty domain string")
	errors.Assert(!cfg.Force, "Expected force to be false.")
	errors.Assert(cfg.Mode.IsNotSet(), "Expected notset mode")
	errors.Assert(cfg.Output == "", "Expected empty output filename.")
	errors.Assert(cfg.Pattern.String() == defaultFilterPattern, "Expected filter pattern not found.")
	errors.Assert(cfg.RecordTypes.String() == defaultDnsRecordTypes, "Expected dns record types not found.")
	errors.Assert(cfg.Timeout == defaultTimeout, "Expected timeout not found.")
}
