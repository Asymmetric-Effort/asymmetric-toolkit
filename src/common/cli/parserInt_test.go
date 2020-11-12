package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestParserInt(t *testing.T) {
	func() {
		fmt.Println("general test")
		min := 1
		value := 2
		strValue := strconv.Itoa(value)
		max := 5
		parser := ParserInt(min, max)

		err, val := parser(&strValue)
		if err != nil {
			panic(err)
		}
		errors.Assert(val.Type == Integer, "Expected Integer type")
		errors.Assert(val.String() == strValue, "unexpected value.")
		errors.Assert(val.Integer() == value, "unexpected int value.")
		defer func() { recover() }()
		errors.Assert(val.Float() != float64(value), "Expected failure")
		errors.Assert(val.Boolean(), "Expected failure")
		errors.Assert(!val.Boolean(), "Expected failure")
		fmt.Println("General Test Passes")
	}()

	go func() { //Lower and Upper bound.
		fmt.Println("upper/lower bound test started")
		min := 1
		max := 5
		parser := ParserInt(min, max)

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
			strValue := strconv.Itoa(i)
			testRun(strValue, i >= min && i <= max)
		}
		fmt.Println("Lower/Upper bound test passes")
	}()

	go func() { //No upper bound
		fmt.Println("no-upper-bound test started")
		parser := ParserInt(1)

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
		testRun("0", false)
		testRun("1", true)
		for i := 1; i < math.MaxInt32; i *= 2 {
			testRun(strconv.Itoa(i), true)
		}
		fmt.Println("No-Upper-bound test passes")
	}()

	func() { //no lower or upper bound
		fmt.Println("unbounded test started")
		parser := ParserInt()

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
		for i := math.MinInt32; i < math.MaxInt32; i *= 2 {
			testRun(strconv.Itoa(i), true)
		}
		fmt.Println("unbounded test passes")
	}()
}
