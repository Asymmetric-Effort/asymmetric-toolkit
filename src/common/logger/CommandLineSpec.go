package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/LogDestination"
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
	logLevel "github.com/combat-diver-foundation/common/logger/log_level"
)

func CommandLineSpecification(config *cli.Configuration) *cli.Specification {
	return &cli.Specification{
		cli.FlagLogDestination: {
			false,
			true,
			config.Enum(
				cli.FlagLogDestination,
				cli.FlagLogDestinationDefault,
				cli.EnumeratedValues{
					LogDestination.StdoutText,
					LogDestination.FileText,
					LogDestination.SyslogText}),
			cli.FlagLogDestinationText,
		},

		cli.FlagLogFile: {
			false,
			true,
			config.String(
				cli.FlagLogFile,
				cli.FlagLogFileDefault,
				utils.RegExDotPlusMan),
			cli.FlagLogFileText,
		},

		cli.FlagLogLevel: {
			false,
			true,
			config.Enum(
				cli.FlagLogLevel,
				cli.FlagLogLevelDefault,
				cli.EnumeratedValues{
					logLevel.DebugText,
					logLevel.InfoText,
					logLevel.WarningText,
					logLevel.ErrorText,
					logLevel.CriticalText}),
			cli.FlagLogLevelText,
		},

		cli.FlagLogServer: {
			false,
			true,
			config.String(
				cli.FlagLogServer,
				cli.FlagLogServerDefault,
				utils.RegExServerString),
			cli.FlagLogServerText,
		},

		//ToDo: Add TLS options for syslog
	}
}
