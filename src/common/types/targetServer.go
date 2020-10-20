package types

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"strconv"
	"strings"
)

type TargetServer string

func (o *TargetServer) isValid() (r bool) {
	// <tcp|udp>:<ipaddr>:<port>
	s := strings.Split(string(*o), ":")
	if len(s) != 3 {
		return false
	}
	return misc.IsTcpUdpProtocol(s[0]) && misc.IsIpAddr(s[1]) && misc.IsPort(s[2])
}

func (o *TargetServer) Set(s string) {
	*o = TargetServer(s)
	errors.Assert(o.isValid(), "TargetServer::Set(): Expected valid Target Server String: protocol:address:port")
}

func (o *TargetServer) Get() (s string) {
	errors.Assert(o.isValid(), "TargetServer::Get(): Expected valid Target Server String: protocol:address:port")
	return string(*o)
}

func (o *TargetServer) String() string {
	errors.Assert(o.isValid(), "TargetServer::String(): Expected valid Target Server String: protocol:address:port")
	return string(*o)
}
func (o *TargetServer) Port() (p int) {
	s := strings.Split(string(*o), ":")
	errors.Assert(len(s) == 3, "TargetServer::Port(): Badly formatted TargetServer string.  "+
		"Expected protocol:address:port")
	return func() int {
		i, err := strconv.Atoi(s[2])
		if err != nil {
			panic("TargetServer::Port(): Invalid port (string) encountered.  " +
				"Expected number.")
		}
		return i
	}()
}
func (o *TargetServer) Addr() string {
	s := strings.Split(string(*o), ":")
	errors.Assert(len(s) == 3, "TargetServer::Addr(): Badly formatted TargetServer string.  "+
		"Expected protocol:address:port")
	errors.Assert(misc.IsIpAddr(s[1]), "TargetServer::Addr(): Expected valid IP address.")
	return string(s[1])
}

func (o *TargetServer) Protocol() string {
	s := strings.Split(string(*o), ":")
	errors.Assert(len(s) == 3, "TargetServer::Protocol(): Badly formatted TargetServer string.  "+
		"Expected protocol:address:port")
	errors.Assert(misc.IsTcpUdpProtocol(s[0]), "TargetServer::IsTcpUdpProtocol(): Expected udp or tcp protocol.")
	return string(s[0])
}
