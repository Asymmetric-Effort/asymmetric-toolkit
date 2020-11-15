package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogger_TapMerge(t *testing.T) {
	var o Logger
	o.tags.global = make(TagTable, 1)

	tagName := "testTagMerge"
	tagId := o.tags.Create(tagName)
	o.tags.AddGlobal(tagId)

	tags := &[]TagId{
		o.tags.Create("tag_A"),
		o.tags.Create("tag_B"),
		o.tags.Create("tag_C"),
		o.tags.Create("tag_D"),
		o.tags.Create("tag_E"),
		o.tags.Create("tag_F"),
	}
	startLen := len(*tags)
	mergedTags := o.tagMerge(tags)
	mergedTagLen := len(mergedTags)
	errors.Assert((mergedTagLen-startLen) == 1, "expected tag list to be one greater.")
}
