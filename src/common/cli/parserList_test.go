package cli

import (
	"fmt"
	"testing"
)

func TestParseList(t *testing.T) {
	func() {
		testRun := func(v string, delimiter string) {
			fmt.Println("Processing...", v)
			parser := ParserList(delimiter)

			err, val := parser(&v)

			if err != nil {
				panic(err)
			}

			expected := fmt.Sprintf("%d%s%s", len(delimiter), delimiter, v)
			if val.Value != expected {
				t.Errorf("Value mismatch: %s", expected)
			}
		}

		tests := []struct {
			Value     string
			delimiter string
		}{
			{
				"A,B,C,D,E",
				",",
			}, {
				"1.2.3.4",
				".",
			}, {
				"C",
				",",
			}, {
				"C:=D:=",
				":=",
			},
		}

		for _, test := range tests {
			testRun(test.Value, test.delimiter)
		}
	}()
}
