package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestParserFloat(t *testing.T) {
	func() {
		fmt.Println("general test")
		min := 1.0
		value := 2.0
		strValue := "2.0"
		max := 5.0
		parser := ParserFloat(min, max)

		err, val := parser(&strValue)
		if err != nil {
			panic(err)
		}
		errors.Assert(val.Type == Float, "Expected Float type")
		errors.Assert(val.String() == strValue, "unexpected value.")
		errors.Assert(val.Float() == value, "unexpected float value.")
		defer func() { recover() }()
		errors.Assert(val.Integer() != int(value), "Expected failure")
		errors.Assert(val.Boolean(), "Expected failure")
		errors.Assert(!val.Boolean(), "Expected failure")
		fmt.Println("General Test Passes")
	}()

	go func() { //Lower and Upper bound.
		fmt.Println("upper/lower bound test started")
		min := 1.0
		max := 5.0
		parser := ParserFloat(min, max)

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
		for i := min - 10; i <= max*2; i++ {
			strValue := strconv.FormatFloat(i,'f',-1,64)
			testRun(strValue, i >= min && i <= max)
		}
		fmt.Println("Lower/Upper bound test passes")
	}()
	go func() { //No upper bound
		fmt.Println("no-upper-bound test started")
		parser := ParserFloat(1)

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
		testRun("0.0", false)
		testRun("1.0", true)
		for i := 1.0; i < math.SmallestNonzeroFloat64; i+=2.0 {
			testRun(strconv.FormatFloat(i,'f',-1,64), true)
		}
		fmt.Println("No-Upper-bound test passes")
	}()
	func() { //no lower or upper bound
		fmt.Println("unbounded test started")
		parser := ParserFloat()

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
		for i := math.SmallestNonzeroFloat64; i < math.MaxFloat64; i*=1.5 {
			testRun(strconv.FormatFloat(i,'f',-1,64), true)
		}
		fmt.Println("unbounded test passes")
	}()
}
