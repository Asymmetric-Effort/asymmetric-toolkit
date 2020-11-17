package cli

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestArgumentFlags(t *testing.T) {
	assert.Error(noFlag == 0, "Expect 0")
	assert.Error(FlagHelp == 1, "Expect 1")
	assert.Error(FlagVersion == 2, "Expect 2")
	assert.Error(FlagDebug == 3, "Expect 3")
	assert.Error(FlagDebug == 4, "Expect 4")
	assert.Error(FlagDebug == 5, "Expect 5")

	assert.Error(FlagDebug == 8, "Expect 8")
	assert.Error(FlagDebug == 9, "Expect 9")

	assert.Error(FlagDebug == 10, "Expect 10")
	assert.Error(FlagDebug == 11, "Expect 11")
	assert.Error(FlagDebug == 12, "Expect 12")
	assert.Error(FlagDebug == 13, "Expect 13")
	assert.Error(FlagDebug == 14, "Expect 14")
	assert.Error(FlagDebug == 15, "Expect 15")

	assert.Error(FlagDebug == 40, "Expect 40")

	assert.Error(FlagDebug == 60, "Expect 60")
	assert.Error(FlagDebug == 61, "Expect 61")
	assert.Error(FlagDebug == 62, "Expect 62")
	assert.Error(FlagDebug == 63, "Expect 63")
	assert.Error(FlagDebug == 64, "Expect 64")
	assert.Error(FlagDebug == 65, "Expect 65")

	assert.Error(FlagDebug == 900, "Expect 900")
}
