package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestLogger_TagDictionary(t *testing.T) {
	var o TagDictionary = make(TagDictionary, 1)
	o[0] = "test"
	assert.Error(o[0] == "test", "expected map result.")
}
