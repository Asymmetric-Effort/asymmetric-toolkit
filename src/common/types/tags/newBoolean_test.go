package tags

import "testing"

func TestTagNewBoolean(t *testing.T){
	var o Boolean = make(Boolean,1)
	o["test"]=true
}