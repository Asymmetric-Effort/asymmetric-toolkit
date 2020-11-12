package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestLogSettings_GetInt(t *testing.T) {
	var o = make(LogSettings,1)
	o["key"]="2"
	assert.Error(o.GetInt("key")==2, "Expected value not found.")
}