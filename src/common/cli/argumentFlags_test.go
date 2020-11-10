package cli

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestArgumentFlags(t *testing.T){
	assert.Error(noFlag == 0, "Expect 0")
	assert.Error(FlagHelp == 1, "Expect 1")
	assert.Error(FlagVersion == 2, "Expect 2")
	assert.Error(FlagDebug == 3, "Expect 3")

}