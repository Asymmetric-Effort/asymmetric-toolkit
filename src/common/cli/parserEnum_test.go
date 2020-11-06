package cli

import (
	"fmt"
	"testing"
)

func TestParserEnum(t *testing.T) {
	parser := ParserEnum("A", "B", "C", "D")

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
			"A",
			true,
		}, {
			"B",
			true,
		}, {
			"C",
			true,
		}, {
			"D",
			true,
		}, {
			"E",
			false,
		},
	}

	for _, test := range tests {
		testRun(test.Value, test.Pass)
	}
}
