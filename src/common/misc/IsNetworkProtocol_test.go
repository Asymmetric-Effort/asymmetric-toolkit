package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"fmt"
	"testing"
)

func TestMiscIsTcpUdpProtocol(t *testing.T) {
	const s0 = "tcp"
	const s1 = "Tcp"
	const s2 = "udp"
	const s3 = "Udp"
	const s4 = "bad"
	errors.Assert(misc.IsNetworkProtocol(s0), fmt.Sprintf("Expected %s", s0))
	errors.Assert(misc.IsNetworkProtocol(s1), fmt.Sprintf("Expected %s", s1))
	errors.Assert(misc.IsNetworkProtocol(s2), fmt.Sprintf("Expected %s", s2))
	errors.Assert(misc.IsNetworkProtocol(s3), fmt.Sprintf("Expected %s", s3))
	errors.Assert(!misc.IsNetworkProtocol(s4), fmt.Sprintf("Expected %s", s4))
}
