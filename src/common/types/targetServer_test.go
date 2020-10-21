package types_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"testing"
)

const (
	udpAny53 = "udp:0.0.0.0:53"
	badAny53 = "bad:0.0.0.0:53"
)

func TestTargetServer(t *testing.T) {
	var ts types.TargetServer = "foo"
	errors.Assert(string(ts) == "foo", "TargetServer expected foo")
}

func TestTargetServer_SetHappy(t *testing.T) {
	var ts types.TargetServer
	errors.Assert(string(ts) == "", "Expected non-empty string")
	ts.Set("udp:127.0.0.1:53")
	errors.Assert(string(ts) == "udp:127.0.0.1:53", "unexpected TargetServer")
	ts.Set("tcp:127.0.0.1:53")
	errors.Assert(string(ts) == "tcp:127.0.0.1:53", "unexpected TargetServer")
	ts.Set("tcp:127.0.0.1:65535")
	errors.Assert(string(ts) == "tcp:127.0.0.1:65535", "unexpected TargetServer")
	ts.Set("tcp:127.0.0.1:0")
	errors.Assert(string(ts) == "tcp:127.0.0.1:0", "unexpected TargetServer")
}

func TestTargetServer_isValid(t *testing.T) {
	var tests = []struct {
		server string
		valid  bool
	}{
		{
			"udp:127.0.0.1:0",
			true,
		},
		{
			"udp:127.0.0.1:53",
			true,
		},
		{
			"udp:127.0.0.1:65535",
			true,
		},
		{
			"udp:127.0.0.1:65536",
			false,
		},
		{
			"tcp:127.0.0.1:53",
			true,
		},
		{
			badAny53,
			false,
		},
		{
			udpAny53,
			true,
		},
		{
			"udp::53",
			false,
		},
		{
			"udp:127.0.0.1:",
			false,
		},
	}
	for i, test := range tests {
		var ts types.TargetServer = types.TargetServer(test.server)
		errors.Assert(ts.IsValid() && test.valid,
			fmt.Sprintf("Test %d: Expected valid: %t, Found valid: %t for string %s",
				i, test.valid, ts.IsValid(), test.server))
	}
}

func TestTargetServer_SetNegPort(t *testing.T) {
	var ts types.TargetServer
	defer func() { recover() }()
	ts.Set("tcp:127.0.0.1:-1")
	t.Error("expected error")
}

func TestTargetServer_SetOverPort(t *testing.T) {
	var ts types.TargetServer
	defer func() { recover() }()
	ts.Set("tcp:127.0.0.1:666666")
	t.Error("expected error")
}
func TestTargetServer_SetZeroAddr(t *testing.T) {
	var ts types.TargetServer
	ts.Set(udpAny53)
	errors.Assert(ts == udpAny53, "Expected "+udpAny53)
}

func TestTargetServer_GetHappy(t *testing.T) {

	var ts types.TargetServer = udpAny53
	errors.Assert(ts.Get() == udpAny53, "Get() failed to get expected server string.")
}

func TestTargetServer_GetBad(t *testing.T) {
	var ts types.TargetServer = "udp:0.0.0.0:-1"
	defer func() { recover() }()
	_ = ts.Get()
	t.Error("expected error")
}

func TestTargetServer_StringHappy(t *testing.T) {
	var ts types.TargetServer = udpAny53
	errors.Assert(ts.String() == udpAny53, "String() failed to get expected server string.")
}

func TestTargetServer_StringBad(t *testing.T) {
	var ts types.TargetServer = "udp:0.0.0.0:-1"
	defer func() { recover() }()
	_ = ts.String()
	t.Error("expected error")
}

func TestTargetServer_Port(t *testing.T) {
	var ts types.TargetServer = udpAny53
	errors.Assert(ts.Port() == 53, "Expected port 53")
}

func TestTargetServer_PortMissing(t *testing.T) {
	var ts types.TargetServer = "udp:0.0.0.0"
	defer func() { recover() }()
	_ = ts.Port()
	t.Error("expect error")
}

func TestTargetServer_PortBad(t *testing.T) {
	var ts types.TargetServer = "udp:0.0.0.0:a"
	defer func() { recover() }()
	_ = ts.Port()
	t.Error("expect error")
}

func TestTargetServer_Address(t *testing.T) {
	var ts types.TargetServer = udpAny53
	errors.Assert(ts.Addr() == "0.0.0.0", "Expected addr 0.0.0.0")
}

func TestTargetServer_AddrMissing(t *testing.T) {
	var ts types.TargetServer = "udp"
	defer func() { recover() }()
	_ = ts.Addr()
	t.Error("expect error")
}

func TestTargetServer_AddrBad(t *testing.T) {
	var ts types.TargetServer = "udp:0000:a"
	defer func() { recover() }()
	_ = ts.Addr()
	t.Error("expect error")
}

func TestTargetServer_Protocol(t *testing.T) {
	var ts types.TargetServer = udpAny53
	errors.Assert(ts.Addr() == "0.0.0.0", "Expected addr 0.0.0.0")
}

func TestTargetServer_ProtocolMissing(t *testing.T) {
	var ts types.TargetServer = "0.0.0.0:53"
	defer func() { recover() }()
	_ = ts.Protocol()
	t.Error("expect error")
}

func TestTargetServer_ProtocolEmpty(t *testing.T) {
	var ts types.TargetServer = ":0.0.0.0:53"
	defer func() { recover() }()
	_ = ts.Protocol()
	t.Error("expect error")
}

func TestTargetServer_ProtocolBad(t *testing.T) {
	var ts types.TargetServer = badAny53
	defer func() { recover() }()
	_ = ts.Protocol()
	t.Error("expect error")
}
