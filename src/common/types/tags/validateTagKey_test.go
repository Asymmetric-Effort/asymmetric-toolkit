package tags

import "testing"

func TestTag_validateTagKey(t *testing.T){
	tests:=[]struct {
		word string
		pass bool
	}{
		{
			"tag",
			true,
		},{
			"tag1",
			true,
		},{
			"ta-g",
			false,
		},{
			"t_ag",
			true,
		},{
			"ta.g",
			true,
		},
	}
	onErrorPanic:=func(){}
	onErrorRecover:=func(){recover()}
	onError:=onErrorPanic
	testFunc:=func(word string, pass bool){
		if pass{
			onError=onErrorPanic
		}else{
			onError=onErrorRecover
		}
		defer onError()
		if !validateTag(word){
			panic("Bad tag key "+word)
		}
	}
	for _,test:=range tests{
		testFunc(test.word,test.pass)
	}
}