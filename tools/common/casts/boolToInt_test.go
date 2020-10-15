package casts

import "testing"

func TestBoolToInt(t *testing.T) {
	_ = BoolToInt(true)
	_ = BoolToInt(false)
}
