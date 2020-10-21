package counter_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/counter"
	"testing"
)

func TestCounter_SetupEmptyCharSet(t *testing.T) {
	var c counter.Counter
	defer func() { _ = recover() }()
	c.Setup("", 5)
	t.Error("Expected error when given empty charset.")
}

func TestCounter_SetupNegativeWordSize(t *testing.T) {
	var c counter.Counter
	defer func() { _ = recover() }()
	c.Setup("0123456789", -1)
	t.Error("Expected error when given negative wordsize.")
}

func TestCounter_SetupZeroWordSize(t *testing.T) {
	var c counter.Counter
	defer func() { _ = recover() }()
	c.Setup("0123456789", 0)
	t.Error("Expected error when given zero (0) wordsize.")
}

func TestCounter_SetupHappy(t *testing.T) {
	const (
		charset  = "0123456789"
		wordsize = 3
	)
	var c counter.Counter
	c.Setup(charset, wordsize)
}
