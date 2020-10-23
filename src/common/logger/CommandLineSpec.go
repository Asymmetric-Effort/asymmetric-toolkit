package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/LogDestination"
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
	logLevel "github.com/combat-diver-foundation/common/logger/log_level"
)

func CommandLineSpecification(config *cli.Configuration) *cli.Specification {
	return &cli.Specification{

		cli.FlagPrefix + FlagDebug: {
			cli.NotRequired,
			cli.NoValueNeeded,
			config.Bool(FlagDebug, false),
			FlagDebugText,
		},

		cli.FlagPrefix + FlagLogDestination: {
			false,
			true,
			config.Enum(
				FlagLogDestination,
				FlagLogDestinationDefault,
				cli.EnumeratedValues{
					LogDestination.StdoutText,
					LogDestination.FileText,
					LogDestination.SyslogText}),
			FlagLogDestinationText,
		},

		cli.FlagPrefix + FlagLogFile: {
			false,
			true,
			config.String(
				FlagLogFile,
				FlagLogFileDefault,
				utils.RegExDotPlusMan),
			FlagLogFileText,
		},

		cli.FlagPrefix + FlagLogLevel: {
			false,
			true,
			config.Enum(
				FlagLogLevel,
				FlagLogLevelDefault,
				cli.EnumeratedValues{
					logLevel.DebugText,
					logLevel.InfoText,
					logLevel.WarningText,
					logLevel.ErrorText,
					logLevel.CriticalText}),
			FlagLogLevelText,
		},

		cli.FlagPrefix + FlagLogServer: {
			false,
			true,
			config.String(
				FlagLogServer,
				FlagLogServerDefault,
				utils.RegExServerString),
			FlagLogServerText,
		},

		//ToDo: Add TLS options for syslog
	}
}
