package cli

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestArgFlags(t *testing.T) {
	errors.Assert(noFlag == 0, "Expected noFlag (default value) to be 0")
	errors.Assert(domainFlag == 1, "Expected 1")
	errors.Assert(modeFlag == 2, "Expected 2")
	errors.Assert(depthFlag == 3, "Expected 3")
	errors.Assert(patternFlag == 4, "Expected 4")
	errors.Assert(outputFlag == 5, "Expected 5")
	errors.Assert(dictionaryFlag == 6, "Expected 6")
	errors.Assert(concurrencyFlag == 7, "Expected 7")
	errors.Assert(timeoutFlag == 8, "Expected 8")
	errors.Assert(delayFlag == 9, "Expected 9")
	errors.Assert(TargetServersFlag == 10, "Expected 10")
	errors.Assert(recordTypesFlag == 11, "Expected 11")
	errors.Assert(debugFlag == 12, "Expected 12")
	errors.Assert(forceFlag == 13, "Expected 13")
	errors.Assert(usageFlag == 14, "Expected 14")
	errors.Assert(versionFlag == 15, "Expected 15")
	errors.Assert(wordSizeFlag == 16, "Expected 16")
	errors.Assert(maxWordCountFlag == 17, "Expected 17")

}
