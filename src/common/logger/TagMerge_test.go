package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestLogger_TapMerge(t *testing.T) {
	var o Logger
	o.tags.global=make(TagTable,1)

	tagName:="testTagMerge"
	tagId:=o.tags.Create(&tagName)
	o.tags.AddGlobal(tagId)

	getaddr:=func(s string)*string{
		return &s
	}

	tags:=&[]TagId{
		o.tags.Create(getaddr("tag_A")),
		o.tags.Create(getaddr("tag_B")),
		o.tags.Create(getaddr("tag_C")),
		o.tags.Create(getaddr("tag_D")),
		o.tags.Create(getaddr("tag_E")),
		o.tags.Create(getaddr("tag_F")),
	}
	startLen:=len(*tags)
	mergedTags:=o.tagMerge(tags)
	mergedTagLen:=len(mergedTags)
	errors.Assert((mergedTagLen - startLen) == 1, "expected tag list to be one greater.")
}
