package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestTagTracker(t *testing.T) {
	var o TagTracker
	o.lock.Lock()
	assert.Error(o.nextTag == 0, "expect 0")
	assert.Error(o.global == nil, "expect nil global")
	assert.Error(o.tagNames == nil, "expect nil tagNames")
	assert.Error(o.tagIds == nil, "expect nil tagId")
	o.lock.Unlock()
}
