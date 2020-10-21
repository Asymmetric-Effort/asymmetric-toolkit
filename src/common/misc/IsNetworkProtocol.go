package misc

import (
	"strings"
)

func IsTcpUdpProtocol(s string) bool {
	switch strings.ToLower(s) {
	case "tcp", "udp":
		return true
	default:
		return false
	}
}