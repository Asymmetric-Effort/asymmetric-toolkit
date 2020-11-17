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
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
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
		Description: Description,
		ProgramName: ProgramName,
		Copyright:   buildconfig.Copyright,
		Version:     buildconfig.Version,
		Argument:    map[string]cli.ArgumentDescriptor{
			"in":{
				cli.FlagInputFile,
				cli.String,
				"",
				"Input flag indicating the source file from which the dictionary " +
					"words will be consumed to create a new dictionary file.",
				cli.ParserString(),
				cli.ExpectValue,
			},
			"out":{
				cli.FlagOutputFile,
				cli.String,
				"",
				"Input flag indicating the output dictionary file which will be created by this tool.",
				cli.ParserString(),
				cli.ExpectValue,
			},
		},
	}
	//
	// Update the Specification with standard parameters
	// using the standardized configuration.
	//
	spec.AddLogLevel(logger.Info)
	//
	// Parse the commandline arguments and in response
	// we expect a boolean (exitProgram) and error object
	// which we will then use to determine if we continue
	// executing or abandon hope.
	//
	exit, err = ui.Parse(&spec, &args)

	var ConfigureLog logger.Configuration
	ConfigureLog.Level.Set(logger.Level(ui.Arguments[cli.FlagLogLevel].Integer()))
	ConfigureLog.Destination.Set(logger.Destination(ui.Arguments[cli.FlagLogDestination].Integer()))
	// ToDo: Pass in settings string.
	//
	// Source Configuration (common/source)
	//
	var SourceConfig source.Configuration
	SourceConfig.Dictionary = ui.Arguments[cli.FlagSourceDictionary].String()
	//
	// General Configuration (common/cli) integration.
	//
	var Config Configuration
	Config.Log = ConfigureLog
	Config.Source = SourceConfig
	//
	// General Log configuration.
	//
	Config.Debug = ui.Arguments[cli.FlagDebug].Boolean()
	Config.Force = ui.Arguments[cli.FlagForce].Boolean()
	Config.InFile = ui.Arguments[cli.FlagInputFile].String()
	Config.OutFile = ui.Arguments[cli.FlagOutputFile].String()



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
