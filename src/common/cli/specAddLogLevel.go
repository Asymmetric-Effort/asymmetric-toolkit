package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
	"fmt"
)

/*
	Specification::AddLogLevel() implements --loglevel <string> flags.
*/
const (
	logLevelHelpText = "Specify the log level. (%s)"
	logLevelArgLong  = "loglevel"
)

func (o *Specification) AddLogLevel(defaultValue logLevel.LogLevel) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[domainArgLong] = ArgumentDescriptor{
		FlagLogLevel,
		String,
		defaultValue.String(),
		fmt.Sprintf(logLevelHelpText, logLevel.LevelStrings),
		ParserEnum(
			logLevel.LogLevel.String(logLevel.Critical),
			logLevel.LogLevel.String(logLevel.Error),
			logLevel.LogLevel.String(logLevel.Warning),
			logLevel.LogLevel.String(logLevel.Info),
			logLevel.LogLevel.String(logLevel.Debug)),
		ExpectValue,
	}
}
