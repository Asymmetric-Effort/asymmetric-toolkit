package tags

import "testing"

func TestTagFloat64(t *testing.T){
	var o Float64 = make(Float64,1)
	o["test"]=1.3
}
