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
	var spec cli.Specification = cli.Specification{
		Author:      buildconfig.Author,
		AuthorEmail: buildconfig.AuthorEmail,
		Description: Description,
		ProgramName: ProgramName,
		Copyright:   buildconfig.Copyright,
		Version:     buildconfig.Version,
		Argument:    map[string]cli.ArgumentDescriptor{},
	}
	//
	// Update the Specification with standard parameters
	// using the standardized configuration.
	//
	spec.AddConcurrency(DefaultConcurrency)
	spec.AddDelay(DefaultAttackDelay)
	spec.AddDepth(DefaultAttackDepth)
	spec.AddSourceDictionary(DefaultSourceDictionary)
	spec.AddDnsRecordType(DnsRecordTypes)
	spec.AddDnsServer(DefaultDnsServer)
	spec.AddDomain(DefaultDnsDomain)
	spec.AddMaxWordCount(DefaultMaxWordSize)
	spec.AddMaxWordSize(DefaultMaxWordSize)
	spec.AddMinWordSize(DefaultMinWordSize)
	spec.AddOutput(DefaultOutput)
	spec.AddSourcePattern(DefaultSourcePattern)
	spec.AddSource(DefaultSource)
	spec.AddTimeout(DefaultTimeout)
	spec.AddLogLevel(logger.Info)
	//
	//ToDo: Integration for the logger to the Specification and common/cli.
	//
	//spec.AddLogDestination("stdout")
	//
	// Parse the commandline arguments and in response
	// we expect a boolean (exitProgram) and error object
	// which we will then use to determine if we continue
	// executing or abandon hope.
	//
	exit, err = ui.Parse(&spec, &args)
	//
	// Log configuration (common/log)
	//
	var ConfigureLog logger.Configuration
	ConfigureLog.Level.Set(logger.Level(ui.Arguments[cli.FlagLogLevel].Integer()))
	ConfigureLog.Destination.Set(logger.Destination(ui.Arguments[cli.FlagLogDestination].Integer()))
	// ToDo: Pass in settings string.
	//
	// Source Configuration (common/source)
	//
	var SourceConfig source.Configuration
	SourceConfig.Dictionary = ui.Arguments[cli.FlagSourceDictionary].String()
	SourceConfig.Mode = ui.Arguments[cli.FlagSource].String()
	SourceConfig.MaxWordCount = ui.Arguments[cli.FlagSourceMaxWordCount].Integer()
	SourceConfig.MaxWordSize = ui.Arguments[cli.FlagSourceMaxWordSize].Integer()
	SourceConfig.MinWordSize = ui.Arguments[cli.FlagSourceMinWordSize].Integer()
	SourceConfig.Pattern = ui.Arguments[cli.FlagSourcePattern].String()
	//
	// General Configuration (common/cli) integration.
	//
	var Config Configuration
	Config.Log = ConfigureLog
	Config.Source = SourceConfig
	//
	// General Log configuration.
	//
	Config.Concurrency = ui.Arguments[cli.FlagConcurrency].Integer()
	Config.Debug = ui.Arguments[cli.FlagDebug].Boolean()
	Config.Delay = ui.Arguments[cli.FlagDelay].Integer()
	Config.Depth = ui.Arguments[cli.FlagDepth].Integer()
	Config.DnsRecordTypes = ui.Arguments[cli.FlagDnsRecordType].String()
	Config.DnsServer = ui.Arguments[cli.FlagDnsServer].String()
	Config.Domain = ui.Arguments[cli.FlagDomain].String()
	Config.Force = ui.Arguments[cli.FlagForce].Boolean()
	Config.Output = ui.Arguments[cli.FlagOutput].String()
	Config.Timeout = ui.Arguments[cli.FlagTimeout].Integer()
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
