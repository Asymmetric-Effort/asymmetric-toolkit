package main

/*
	Define and process the commandline specification using standardized
	arguments then extract the result into the Configuration struct and return
	the same by reference.  When this function falls out of scope the memory
	used by the commandline specification will be freed.
*/

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/buildConfig"
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"fmt"
)

func ProcessSpecification(args []string) (cfg *Configuration, exit bool, err error) {
	//
	// The cli.CommandLine struct is the top-level
	// for the commandline interface which will be
	// loaded with a parsed and processed command-
	// line parameter set.
	//
	var ui cli.CommandLine
	//
	// The Specification defines the commandline
	// parameters this program will use when it
	// is executed.
	//
	var spec = cli.Specification{
		Author:      buildconfig.Author,
		AuthorEmail: buildconfig.AuthorEmail,
		Description: ProgramDescription,
		ProgramName: ProgramName,
		Copyright:   buildconfig.Copyright,
		Version:     buildconfig.Version,
		Argument:    map[string]cli.ArgumentDescriptor{},
	}
	//
	// Update the Specification with standard parameters
	// using the standardized configuration.
	//
	spec.AddInputFile()
	spec.AddOutputFile()
	//
	// Parse the commandline arguments and in response
	// we expect a boolean (exitProgram) and error object
	// which we will then use to determine if we continue
	// executing or abandon hope.
	//
	exit, err = ui.Parse(&spec, &args)
	if err != nil {
		return
	}
	//
	// Configure common/logger
	//
	var CfgLog logger.Configuration
	CfgLog.Level.SetString(ui.Arguments[cli.FlagLogLevel].String())
	CfgLog.Destination.SetString(ui.Arguments[cli.FlagLogDestination].String())
	// ToDo: Pass in settings string.
	//
	// General Configuration (common/cli) integration.
	//
	var Config Configuration
	Config.Log = CfgLog
	//
	// General Log configuration.
	//
	Config.Debug = ui.Arguments[cli.FlagDebug].Boolean()          // Add debug logging
	Config.Force = ui.Arguments[cli.FlagForce].Boolean()          // Add force to overwrite targets
	Config.InputFile = ui.Arguments[cli.FlagInputFile].String()   // Specify the input file (text file)
	Config.OutputFile = ui.Arguments[cli.FlagOutputFile].String() // Specify the output file (dictionary)
	//
	// Evaluate the error object.
	//
	if err != nil {
		cfg = nil
		err = fmt.Errorf("error:%v", err)
		exit = true
		return
	}
	cfg = &Config
	return
}
