package tags

import "testing"

func TestNewTag(t *testing.T) {
	var tag Tag = NewTag()
	tag["foo"] = struct{}{}
}
