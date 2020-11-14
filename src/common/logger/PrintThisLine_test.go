package logger

import (
	assert "github.com/sam-caldwell/adrestia-assertions/src"
	"testing"
)

func TestLogger_PrintThisLine(t *testing.T) {
	var o Logger
	o.Level = Debug
	assert.Error(o.PrintThisLine(Debug), "expected true")
	assert.Error(o.PrintThisLine(Info), "expected true")
	assert.Error(o.PrintThisLine(Warning), "expected true")
	assert.Error(o.PrintThisLine(Error), "expected true")
	assert.Error(o.PrintThisLine(Critical), "expected true")

	o.Level = Info
	assert.Error(!o.PrintThisLine(Debug), "expected false")
	assert.Error(o.PrintThisLine(Info), "expected true")
	assert.Error(o.PrintThisLine(Warning), "expected true")
	assert.Error(o.PrintThisLine(Error), "expected true")
	assert.Error(o.PrintThisLine(Critical), "expected true")

	o.Level = Warning
	assert.Error(!o.PrintThisLine(Debug), "expected false")
	assert.Error(!o.PrintThisLine(Info), "expected false")
	assert.Error(o.PrintThisLine(Warning), "expected true")
	assert.Error(o.PrintThisLine(Error), "expected true")
	assert.Error(o.PrintThisLine(Critical), "expected true")

	o.Level = Error
	assert.Error(!o.PrintThisLine(Debug), "expected false")
	assert.Error(!o.PrintThisLine(Info), "expected false")
	assert.Error(!o.PrintThisLine(Warning), "expected false")
	assert.Error(o.PrintThisLine(Error), "expected true")
	assert.Error(o.PrintThisLine(Critical), "expected true")

	o.Level = Critical
	assert.Error(!o.PrintThisLine(Debug), "expected false")
	assert.Error(!o.PrintThisLine(Info), "expected false")
	assert.Error(!o.PrintThisLine(Warning), "expected false")
	assert.Error(!o.PrintThisLine(Error), "expected false")
	assert.Error(o.PrintThisLine(Critical), "expected true")

}
