package main_test

/*
	This is the top-level test program for testing CliDev.
 */

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tests := []struct {
		name string
		code int
		args []string
	}{
		{
			name: "no args",
			code: 0,
			args: []string{""},
		},
		{
			name: "short help",
			code: 0,
			args: []string{
				"-h",
			},
		}, {
			name: "long help:,",
			code: 0,
			args: []string{
				"--help",
			},
		}, {
			name: "short version",
			code: 0,
			args: []string{
				"-v",
			},
		}, {
			name: "long version",
			code: 0,
			args: []string{
				"--version",
			},
		}, {
			name: "bad flag (expect error)",
			code: 1,
			args: []string{
				"--badFlagShouldErr",
			},
		},
	}
	startArgs := os.Args //Get state
	for _, test := range tests {
		os.Args = startArgs //Reset args to original state.
		for _, i := range test.args {
			//Copy in new args
			os.Args = append(os.Args, i)
		}
		fmt.Printf("Starting test: %s (%v) code:%d\n", test.name, test.args, test.code)
		if m.Run() == test.code {
			fmt.Printf("TestMain() exit as expected with code %d\n\n", test.code)
			continue
		} else {
			fmt.Println("Unknown or unexpected error")
			os.Exit(1)
		}
	}
}
