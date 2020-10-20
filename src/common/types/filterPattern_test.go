package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestFilterPattern_Set_Happy(t *testing.T) {
	var fp FilterPattern
	fp.Set(`^a$`)
	errors.Assert(fp.re.MatchString("a"), "Failed positive match test")
	errors.Assert(!fp.re.MatchString("b"), "Failed negative match test")
}

func TestFilterPattern_Set_Sad(t *testing.T) {
	var fp FilterPattern
	defer func() { recover() }()
	fp.Set("]BadRegex[^")
	t.Fatal("Bad Regular expression should have caused error.")
}

func TestFilterPattern_Match(t *testing.T) {
	var fp FilterPattern
	fp.Set(`^a$`)
	errors.Assert(fp.Match("a"), "Failed positive match test")
	errors.Assert(!fp.Match("b"), "Failed negative match test")
}

func TestFilterPattern_String(t *testing.T) {
	var fp FilterPattern
	pattern:=`^a$`
	fp.Set(pattern)
	errors.Assert(pattern == fp.String(), "Expected pattern string mismatch")
}