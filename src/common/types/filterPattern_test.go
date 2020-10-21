package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestFilterPattern_Set_Happy(t *testing.T) {
	var fp types.FilterPattern
	fp.Set(`^a$`)
	errors.Assert(fp.Re.MatchString("a"), "Failed positive match test")
	errors.Assert(!fp.Re.MatchString("b"), "Failed negative match test")
}

func TestFilterPattern_Set_Sad(t *testing.T) {
	var fp types.FilterPattern
	defer func() { recover() }()
	fp.Set("]BadRegex[^")
	t.Fatal("Bad Regular expression should have caused error.")
}

func TestFilterPattern_Match(t *testing.T) {
	var fp types.FilterPattern
	fp.Set(`^a$`)
	errors.Assert(fp.Match("a"), "Failed positive match test")
	errors.Assert(!fp.Match("b"), "Failed negative match test")
}

func TestFilterPattern_String(t *testing.T) {
	var fp types.FilterPattern
	pattern:=`^a$`
	fp.Set(pattern)
	errors.Assert(pattern == fp.String(), "Expected pattern string mismatch")
}