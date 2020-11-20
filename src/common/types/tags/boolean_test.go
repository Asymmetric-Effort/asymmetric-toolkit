package tags

import "testing"

func TestTagBoolean(t *testing.T){
	var o Boolean = make(Boolean,1)
	o["test"]=true
}