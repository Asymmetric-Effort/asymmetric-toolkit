package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/destination"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/level"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
)

type Configuration struct {
	Concurrency   types.PositiveInteger
	Debug         bool
	Delay         types.PositiveInteger
	Depth         types.PositiveInteger
	Dictionary    types.FilePath
	TargetServer  types.TargetServer
	Domain        types.DomainName
	Force         bool
	Log           struct {
		Destination destination.LogDestination
		Target      string //Could be log file or syslog server.
		Level       level.LogLevel
	}
	Mode         types.DataSource
	Output       types.FilePath
	Pattern      types.FilterPattern
	RecordTypes  types.List
	Timeout      types.PositiveInteger
	WordSize     types.PositiveInteger
	MaxWordCount types.PositiveInteger
}
