package cli

/*
	Specification::AddLogDestination() implements the --logout <filename>.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"fmt"
)

const (
	logDestinationHelpText string = "Log destination (out) flag indicating the output channel to " +
		"which logs will be written (%s)"
	logDestinationArgLong = "logdestination"
)

func (o *Specification) AddLogDestination(defaultValue logger.Destination) {
	//
	// We add a long argument for
	//
	o.Initialize()
	o.Argument[logDestinationArgLong] = ArgumentDescriptor{
		FlagLogDestination,
		String,
		defaultValue.String(),
		fmt.Sprintf(logDestinationHelpText, logger.DestinationStrings),
		ParserEnum("stdout","file","syslog"),
		ExpectValue,
	}
}
