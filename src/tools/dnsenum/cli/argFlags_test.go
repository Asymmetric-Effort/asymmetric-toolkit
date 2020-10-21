package cli_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"testing"
)

func TestArgFlags(t *testing.T) {
	errors.Assert(cli.NoFlag == 0, "Expected noFlag (default value) to be 0")
	errors.Assert(cli.DomainFlag == 1, "Expected 1")
	errors.Assert(cli.ModeFlag == 2, "Expected 2")
	errors.Assert(cli.DepthFlag == 3, "Expected 3")
	errors.Assert(cli.PatternFlag == 4, "Expected 4")
	errors.Assert(cli.OutputFlag == 5, "Expected 5")
	errors.Assert(cli.DictionaryFlag == 6, "Expected 6")
	errors.Assert(cli.ConcurrencyFlag == 7, "Expected 7")
	errors.Assert(cli.TimeoutFlag == 8, "Expected 8")
	errors.Assert(cli.DelayFlag == 9, "Expected 9")
	errors.Assert(cli.TargetServerFlag == 10, "Expected 10")
	errors.Assert(cli.RecordTypesFlag == 11, "Expected 11")
	errors.Assert(cli.DebugFlag == 12, "Expected 12")
	errors.Assert(cli.ForceFlag == 13, "Expected 13")
	errors.Assert(cli.UsageFlag == 14, "Expected 14")
	errors.Assert(cli.VersionFlag == 15, "Expected 15")
	errors.Assert(cli.WordSizeFlag == 16, "Expected 16")
	errors.Assert(cli.MaxWordCountFlag == 17, "Expected 17")

}
