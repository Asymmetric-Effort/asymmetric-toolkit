package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"strconv"
	"testing"
)

func TestPositiveInteger_Set_Happy(t *testing.T) {
	var p PositiveInteger
	validSet := []string{"1", "2", "3", "4", "5", "255", "1048576", "1073741824"}
	for _, v := range validSet {
		p.Set(v)
		i, e := strconv.Atoi(v)
		if e != nil {
			t.Fatalf("Error parsing %s into integer.  Error: %v", v, e)
		}
		errors.Assert(p == PositiveInteger(i), "Conversion error in PositiveInteger type")
	}
}

func TestPositiveInteger_Set_Zero(t *testing.T) {
	var p PositiveInteger
	var v string = "0"
	defer func() { recover() }()
	p.Set(v)
	t.Error("Expected failure given invalid zero value PositiveInteger.")

}

func TestPositiveInteger_Set_Negative(t *testing.T) {
	var p PositiveInteger
	var v string = "-1"
	defer func() { recover() }()
	p.Set(v)
	t.Error("Expected failure given invalid negative integer in PositiveInteger.")
}

func TestPositiveInteger_Set_Empty(t *testing.T) {
	var p PositiveInteger
	var v string = ""
	defer func() { recover() }()
	p.Set(v)
	t.Error("Expected failure given invalid empty string.")
}
