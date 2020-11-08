package cli

import (
	"fmt"
	"strings"
	"testing"
)

func TestArgumentType_String(t *testing.T) {
	tests := []struct {
		token    ArgumentType
		expected string
		pass     bool
	}{
		{
			-1,
			"Bad1",
			false,
		}, {
			None,
			"None",
			true,
		}, {
			String,
			"String",
			true,
		}, {
			Integer,
			"Integer",
			true,
		}, {
			Float,
			"Float",
			true,
		}, {
			Boolean,
			"Boolean",
			true,
		}, {
			5,
			"Bad2",
			false,
		},
	}
	noop := func() {
		fmt.Println("Error occurred but was NOT recovered.")
	}
	recoverErrorFunc := noop

	fmt.Println("---")
	testRun := func(pass bool, token ArgumentType, expected string) {

		fmt.Printf("Testing... (%d): '%s' (%v)\n", token, expected, pass)
		if pass {
			fmt.Printf("Expect pass\n")
			recoverErrorFunc = noop
		} else {
			fmt.Println("Expect failure")
			recoverErrorFunc = func() {
				fmt.Println("Error occurred but was recovered")
				recover()
			}
		}

		defer recoverErrorFunc()

		fmt.Printf("Preparing to call .String() method for %d\n",token)
		if token.String() == expected {
			fmt.Println("We're ok")
			return
		} else {
			panic(fmt.Sprintf("Token(%d) failed: %s", token, expected))
		}
	}

	fmt.Println("lenArgTypes:", len(strings.Split(argumentTypes, comma)))
	for i, test := range tests {
		fmt.Println("Start: ", i)
		testRun(test.pass, test.token, test.expected)
		fmt.Println("Pass: ", i)
		fmt.Println("---")
	}
}
