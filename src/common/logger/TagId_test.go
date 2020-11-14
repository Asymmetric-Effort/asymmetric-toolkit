package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"math"
	"testing"
)

func TestTag(t *testing.T) {
	func() {
		var tag TagId = 0
		assert.Error(tag>=0,"tag greater than zero")

		tag = math.MaxUint32
		assert.Error(tag<=math.MaxUint32, "tag less than 2^32")
	}()
}
