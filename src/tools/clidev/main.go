package main

/*
	CliDev
	------
	This is the commandline development program intended for use in
	developing and testing the common/cli module and at the same time
	providing a template for how it is used by programmers working to
	build new tools in the asymmetric-toolkit.
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"fmt"
	"os"
)

const (
	DefaultAttackDelay      int    = 1                    // Default number of seconds between attack payload executions.
	DefaultConcurrency      int    = 1                    // Default number of concurrent queries to run.
	DefaultAttackDepth      int    = 20                   // Default number of DNS subdomain levels to attack
	DefaultDnsDomain        string = "localdomain"        // Default domain to be enumerated.
	DefaultDnsServer        string = "udp://127.0.0.1:53" // Default DNS server.
	DefaultTimeout          int    = 60                   // Default number of seconds before connection timeout.
	DefaultMinWordSize      int    = 3                    // Default minimum character length of each word
	DefaultMaxWordSize      int    = 5                    // Default maximum character length of each word.
	DefaultOutput           string = "output.txt"         // Default output path/file for reporting results.
	DefaultSource           string = "sequence"           // Default word source for our enumeration attack.
	DefaultSourceDictionary string = ""                   //Default path/file to a dictionary (we will have none).
	DefaultSourcePattern    string = ".+"                 // Default regex pattern string to filter words to be used in an attack.
	DnsRecordTypes                 = "A,AAAA,MX,TXT,SOA,NS,CNAME"
)

type Configuration struct {
	Concurrency    int
	delay          int
	depth          int
	dictionary     string
	dnsServer      string
	dnsRecordTypes string
	domain         string
	maxWordCount   int
	minWordSize    int
	maxWordSize    int
	output         string
	pattern        string
	source         string
	timeout        int
}

func main() {
	fmt.Println("Setup Specification...")
	var spec cli.Specification = cli.Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		Description: "Test program for commandline interface.",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument: map[string]cli.ArgumentDescriptor{
			"test": {
				1000, // >= 1000 is a project-defined FlagId
				cli.None,
				"noop default won't change",
				"This is an option.",
				cli.ParserNoop(),
				cli.ExpectFlag,
			},
			"myUnboundedIntOption": {
				1001, // >= 1000 is a project-defined FlagId
				cli.Integer,
				"5",
				"This is a Second Option",
				cli.ParserInt(),
				cli.ExpectValue,
			},
			"myIntOption": {
				1001, // >= 1000 is a project-defined FlagId
				cli.Integer,
				"5",
				"This is a Second Option",
				cli.ParserInt(0, 10),
				cli.ExpectValue,
			},
			"myUnboundedStringOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"test",
				"This is a Second Option",
				cli.ParserString(),
				cli.ExpectValue,
			},
			"myStringOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"test",
				"This is a Second Option",
				cli.ParserString(`.+`),
				cli.ExpectValue,
			},
			"myUnboundedFloatOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"1.2",
				"This is a Second Option",
				cli.ParserFloat(),
				cli.ExpectValue,
			},
			"myFloatOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"1.2",
				"This is a Second Option",
				cli.ParserFloat(0.0, 2.0),
				cli.ExpectValue,
			},
			"myBooleanOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"false",
				"This is a Second Option",
				cli.ParserBool(),
				cli.ExpectValue,
			},
			"myEnumOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"a",
				"This is a Second Option",
				cli.ParserEnum("a", "b", "c"),
				cli.ExpectValue,
			},
			"myListOption": {
				1002, // >= 1000 is a project-defined FlagId
				cli.String,
				"a",
				"This is a Second Option",
				cli.ParserList(","),
				cli.ExpectValue,
			},
		},
	}
	spec.AddConcurrency(DefaultConcurrency)
	spec.AddDelay(DefaultAttackDelay)
	spec.AddDepth(DefaultAttackDepth)
	spec.AddSourceDictionary(DefaultSourceDictionary)

	spec.AddDnsRecordType(DnsRecordTypes)
	spec.AddDnsServer(DefaultDnsServer)
	spec.AddDomain(DefaultDnsDomain)
	spec.AddMinWordSize(DefaultMinWordSize)
	spec.AddMaxWordSize(DefaultMaxWordSize)
	spec.AddMaxWordCount(DefaultMaxWordSize)
	spec.AddOutput(DefaultOutput)
	spec.AddSourcePattern(DefaultSourcePattern)
	spec.AddSource(DefaultSource)
	spec.AddTimeout(DefaultTimeout)

	fmt.Println("Setup CommandLine struct...")
	var ui cli.CommandLine
	fmt.Println("Parse Arguments...")
	args := []string{"--concurrency", "999", "--depth", "2"}
	exit, err := ui.Parse(&spec, &args)
	fmt.Println("Evaluate results...")

	var Config Configuration
	fmt.Println("===")
	fmt.Println("Flag:", cli.FlagConcurrency)
	fmt.Println("Argument:", ui.Arguments)
	fmt.Println("===")
	Config.Concurrency = ui.Arguments[cli.FlagConcurrency].Integer()
	Config.delay = ui.Arguments[cli.FlagDelay].Integer()
	Config.depth = ui.Arguments[cli.FlagDepth].Integer()
	Config.dictionary = ui.Arguments[cli.FlagSourceDictionary].String()
	Config.dnsRecordTypes = ui.Arguments[cli.FlagDnsRecordType].String()
	Config.dnsServer = ui.Arguments[cli.FlagDnsServer].String()
	Config.domain = ui.Arguments[cli.FlagDomain].String()
	Config.maxWordCount = ui.Arguments[cli.FlagSourceMaxWordCount].Integer()
	Config.minWordSize = ui.Arguments[cli.FlagSourceMinWordSize].Integer()
	Config.maxWordSize = ui.Arguments[cli.FlagSourceMaxWordSize].Integer()
	Config.output = ui.Arguments[cli.FlagOutput].String()
	Config.pattern = ui.Arguments[cli.FlagSourcePattern].String()
	Config.source = ui.Arguments[cli.FlagSource].String()
	Config.timeout = ui.Arguments[cli.FlagTimeout].Integer()
	if err != nil {
		fmt.Printf("Error:%v", err)
		os.Exit(cli.ErrArgumentParseError)
	}
	if exit {
		fmt.Println("Success")
		os.Exit(cli.ErrSuccess)
	}
}
