package cli

import (
	"fmt"
	"testing"
)

func TestParserString(t *testing.T) {
	func() {
		parser := ParserString("[a-z]+")

		testRun := func(v string, p bool) {
			if !p {
				fmt.Printf("Expect failure for %v\n", v)
				defer func() { recover() }()
			}
			err, val := parser(&v)
			if err != nil {
				panic(err)
			}
			if val.Value != v {
				t.Errorf("Value mismatch: %s", v)
			}
		}

		tests := []struct {
			Value string
			Pass  bool
		}{
			{
				"apassingtest",
				true,
			}, {
				"Afailedtest",
				false,
			}, {
				"Another failed test",
				false,
			}, {
				"Numeric failed test 123",
				true,
			},
		}

		for _, test := range tests {
			testRun(test.Value, test.Pass)
		}
	}()
	func() {
		parser := ParserString("[a-zA-Z0-9 ]+")

		testRun := func(v string, p bool) {
			if !p {
				fmt.Printf("Expect failure for %v\n", v)
				defer func() { recover() }()
			}
			err, val := parser(&v)
			if err != nil {
				panic(err)
			}
			if val.Value != v {
				t.Errorf("Value mismatch: %s", v)
			}
		}

		tests := []struct {
			Value string
			Pass  bool
		}{
			{
				"apassingtest",
				true,
			}, {
				"Afailedtest",
				true,
			}, {
				"Another failed test",
				true,
			}, {
				"Numeric failed test 123",
				true,
			},
		}

		for _, test := range tests {
			testRun(test.Value, test.Pass)
		}
	}()
}
