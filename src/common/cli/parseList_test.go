package cli

import (
	"testing"
)

func TestParseList(t *testing.T) {
	func() {
		testRun := func(v string, delimiter string) {

			parser := ParseList(delimiter)

			err, val := parser(&v)

			if err != nil {
				panic(err)
			}

			if val.Value != v {
				t.Errorf("Value mismatch: %s", v)
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
			},
		}

		for _, test := range tests {
			testRun(test.Value, test.delimiter)
		}
	}()
}
