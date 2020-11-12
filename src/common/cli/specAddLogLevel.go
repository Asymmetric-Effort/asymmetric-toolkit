package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"fmt"
)

/*
	Specification::AddLogLevel() implements --loglevel <string> flags.
*/
const (
	logLevelHelpText = "Specify the log level. (%s)"
	logLevelArgLong  = "loglevel"
)

func (o *Specification) AddLogLevel(defaultValue logger.LogLevel) {
	//
	// Initialize the Argument object.
	//
	o.Initialize()
	o.Argument[logLevelArgLong] = ArgumentDescriptor{
		FlagLogLevel,
		String,
		defaultValue.String(),
		fmt.Sprintf(logLevelHelpText, logger.LevelStrings),
		ParserEnum("critical", "error", "warning", "info", "debug"),
		ExpectValue,
	}
}
