package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"strconv"
	"testing"
)

func TestPositiveInteger_Set_Happy(t *testing.T) {
	var p types.PositiveInteger
	validSet := []string{"1", "2", "3", "4", "5", "255", "1048576", "1073741824"}
	for _, v := range validSet {
		p.Set(v)
		i, e := strconv.Atoi(v)
		if e != nil {
			t.Fatalf("Errorf parsing %s into integer.  Errorf: %v", v, e)
		}
		errors.Assert(p == types.PositiveInteger(i), "Conversion error in PositiveInteger type")
	}
}

func TestPositiveInteger_Set_Zero(t *testing.T) {
	var p types.PositiveInteger
	var v string = "0"
	defer func() { recover() }()
	p.Set(v)
	t.Error("Expected failure given invalid zero value PositiveInteger.")

}

func TestPositiveInteger_Set_Negative(t *testing.T) {
	var p types.PositiveInteger
	var v string = "-1"
	defer func() { recover() }()
	p.Set(v)
	t.Error("Expected failure given invalid negative integer in PositiveInteger.")
}

func TestPositiveInteger_Set_Empty(t *testing.T) {
	var p types.PositiveInteger
	var v string = ""
	defer func() { recover() }()
	p.Set(v)
	t.Error("Expected failure given invalid empty string.")
}
