package loggercli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/LogDestination"
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
	logLevel "github.com/combat-diver-foundation/common/logger/log_level"
)

func CommandLineSpecification(config *cli.CommandLine) *cli.Specification {
	return &cli.Specification{

		FlagPrefix + FlagDebug: {
			NotRequired,
			NoValueNeeded,
			config.ValidateBool(FlagDebug, false),
			FlagDebugText,
		},

		FlagPrefix + FlagLogDestination: {
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

		FlagPrefix + FlagLogFile: {
			false,
			true,
			config.ValidateString(
				FlagLogFile,
				FlagLogFileDefault,
				utils.RegExDotPlusMan),
			FlagLogFileText,
		},

		FlagPrefix + FlagLogLevel: {
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

		FlagPrefix + FlagLogServer: {
			false,
			true,
			config.ValidateString(
				FlagLogServer,
				FlagLogServerDefault,
				utils.RegExServerString),
			FlagLogServerText,
		},

		//ToDo: Add TLS options for syslog
	}
}
