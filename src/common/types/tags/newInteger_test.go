package tags

import "testing"

func TestTagNewInteger(t *testing.T){
	var o Integer = make(Integer,1)
	o["test"]=1
}