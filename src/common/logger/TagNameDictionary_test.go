package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestLogger_TagNameDictionary(t *testing.T) {
	var o TagNameDictionary = make(TagNameDictionary, 1)
	o["test"] = 0
	assert.Error(o["test"] == 0, "expected map result.")
}