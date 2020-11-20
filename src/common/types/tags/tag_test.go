package tags

import (
	"testing"
)

func TestTag(t *testing.T){
	var o Tag = make(Tag,1)
	o["test"]=struct{}{}
}