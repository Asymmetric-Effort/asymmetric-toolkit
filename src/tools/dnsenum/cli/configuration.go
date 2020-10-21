package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
)

type Configuration struct {
	Debug        bool
	Force        bool
	reserved1	 [6]byte //padding needed for proper memory alignment
	Concurrency  types.PositiveInteger
	Timeout      types.PositiveInteger
	WordSize     types.PositiveInteger
	MaxWordCount types.PositiveInteger
	Delay        types.PositiveInteger
	Depth        types.PositiveInteger
	reserved2	 [16]byte //padding needed for proper memory alignment
	Dictionary   types.FilePath
	TargetServer types.TargetServer
	Domain       types.DomainName
	Log          struct {
		Destination destination.LogDestination
		Level       level.Level
		Target      string //Could be log file or syslog server.
	}
	Mode        sourcetype.DataSource
	Output      types.FilePath
	Pattern     types.FilterPattern
	RecordTypes types.List
}
