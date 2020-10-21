package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"testing"
)

func TestFilterPattern_Set(t *testing.T) {
	var fp types.FilterPattern
	pattern:=`^a$`
	fp.Set(pattern)
	errors.Assert(fp.String()==pattern, "expected set patttern to be returned.")
	
	//Sad path
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
	pattern := `^a$`
	fp.Set(pattern)
	errors.Assert(pattern == fp.String(), "Expected pattern string mismatch")
}

func TestFilterPattern_Get(t *testing.T) {
	var fp types.FilterPattern
	errors.Assert(fp.Get() == "", "initial state is empty")
	fp.Set(".+")
	errors.Assert(fp.Get() == ".+", "expected .+")
}

func TestFilterPattern_IsNil(t *testing.T) {
	var fp types.FilterPattern
	errors.Assert(fp.IsNil(), "Expected true")
	fp.Set(".+")
	errors.Assert(!fp.IsNil(), "Expected false")
}
