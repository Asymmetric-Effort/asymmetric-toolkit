package cli_test

import (
	cli2 "asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestArgFlags(t *testing.T) {
	errors.Assert(cli2.NoFlag == 0, "Expected noFlag (default value) to be 0")
	errors.Assert(cli2.DomainFlag == 1, "Expected 1")
	errors.Assert(cli2.ModeFlag == 2, "Expected 2")
	errors.Assert(cli2.DepthFlag == 3, "Expected 3")
	errors.Assert(cli2.PatternFlag == 4, "Expected 4")
	errors.Assert(cli2.OutputFlag == 5, "Expected 5")
	errors.Assert(cli2.DictionaryFlag == 6, "Expected 6")
	errors.Assert(cli2.ConcurrencyFlag == 7, "Expected 7")
	errors.Assert(cli2.TimeoutFlag == 8, "Expected 8")
	errors.Assert(cli2.DelayFlag == 9, "Expected 9")
	errors.Assert(cli2.TargetServerFlag == 10, "Expected 10")
	errors.Assert(cli2.RecordTypesFlag == 11, "Expected 11")
	errors.Assert(cli2.DebugFlag == 12, "Expected 12")
	errors.Assert(cli2.ForceFlag == 13, "Expected 13")
	errors.Assert(cli2.UsageFlag == 14, "Expected 14")
	errors.Assert(cli2.VersionFlag == 15, "Expected 15")
	errors.Assert(cli2.WordSizeFlag == 16, "Expected 16")
	errors.Assert(cli2.MaxWordCountFlag == 17, "Expected 17")

}
