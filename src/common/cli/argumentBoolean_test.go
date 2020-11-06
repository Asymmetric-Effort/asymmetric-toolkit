package cli

import (
	"fmt"
	"testing"
)

func TestArgument_Boolean(t *testing.T) {
	tests:=[]struct{
		Value string
		Result bool
		Error bool
	}{
		{
			"true",
			true,
			false,

		},
		{
			"True",
			true,
			false,

		},
		{
			"True ",
			true,
			true,

		},
		{
			" True",
			true,
			true,

		},
		{
			"false",
			false,
			false,

		},
		{
			"False",
			false,
			false,

		},
		{
			" false",
			false,
			true,

		},
		{
			"false ",
			false,
			true,

		},
		{
			"bad",
			false,
			true,

		},
		{
			"0",
			false,
			true,

		},
		{
			"1",
			false,
			true,
		},
	}
	testRun:=func(v string, expectFail bool){
		var o Argument
		if expectFail {
			defer func(){recover()}() //Recover from an error for sad-path tests.
		}
		o.Type = Boolean
		o.Value = v
		b := o.Boolean()
		fmt.Println(b)
	}
	for _,test:=range tests {
		testRun(test.Value, test.Error)
	}
}
