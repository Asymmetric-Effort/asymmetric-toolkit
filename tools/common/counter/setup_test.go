package counter

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"testing"
)

func TestCounter_SetupEmptyCharSet(t *testing.T) {
	var c Counter
	defer func() { recover() }()
	c.Setup("", 5)
	t.Error("Expected error when given empty charset.")
}

func TestCounter_SetupNegativeWordSize(t *testing.T){
	var c Counter
	defer func() { recover() }()
	c.Setup("0123456789", -1)
	t.Error("Expected error when given negative wordsize.")
}

func TestCounter_SetupZeroWordSize(t *testing.T){
	var c Counter
	defer func() { recover() }()
	c.Setup("0123456789", 0)
	t.Error("Expected error when given zero (0) wordsize.")
}

func TestCounter_SetupHappy(t *testing.T){
	var c Counter
	const (
		charset="0123456789"
		wordsize = 3
	)

	c.Setup(charset,wordsize)
	expectedRunes:=[]rune(charset)
	errors.Assert(string(*c.runes)==string(expectedRunes), "Expected runes mismatched.")
	errors.Assert(int(c.maxBase) == len(charset), "Expected wordsize mismatch")
	for i:=0;i<wordsize;i++{
		errors.Assert((*c.data)[i]==0, "Expected counter.data to be initialized with zero state")
	}
}