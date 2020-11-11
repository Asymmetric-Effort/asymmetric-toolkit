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
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger/logLevel"
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
	spec.AddLogLevel(logLevel.Info)
	//
	// Parse the commandline arguments and in response
	// we expect a boolean (exitProgram) and error object
	// which we will then use to determine if we continue
	// executing or abandon hope.
	//
	exit, err = ui.Parse(&spec, &args)
	var Config Configuration
	Config.Concurrency = ui.Arguments[cli.FlagConcurrency].Integer()
	Config.delay = ui.Arguments[cli.FlagDelay].Integer()
	Config.depth = ui.Arguments[cli.FlagDepth].Integer()
	Config.dictionary = ui.Arguments[cli.FlagSourceDictionary].String()
	Config.dnsRecordTypes = ui.Arguments[cli.FlagDnsRecordType].String()
	Config.dnsServer = ui.Arguments[cli.FlagDnsServer].String()
	Config.domain = ui.Arguments[cli.FlagDomain].String()
	Config.maxWordCount = ui.Arguments[cli.FlagSourceMaxWordCount].Integer()
	Config.maxWordSize = ui.Arguments[cli.FlagSourceMaxWordSize].Integer()
	Config.minWordSize = ui.Arguments[cli.FlagSourceMinWordSize].Integer()
	Config.output = ui.Arguments[cli.FlagOutput].String()
	Config.pattern = ui.Arguments[cli.FlagSourcePattern].String()
	Config.source = ui.Arguments[cli.FlagSource].String()
	Config.timeout = ui.Arguments[cli.FlagTimeout].Integer()
	Config.log.level = logLevel.LogLevel(ui.Arguments[cli.FlagLogLevel].Integer())
	Config.log.facility = LogFacility.Facility(ui.Arguments[cli.FlagLogFacility].String())
	Config.log.target = ui.Arguments[cli.FlagTarget]
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
