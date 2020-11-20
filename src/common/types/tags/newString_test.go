package tags

import "testing"

func TestTagNewString(t *testing.T){
	var o String = make(String,1)
	o["test"]="true"
}