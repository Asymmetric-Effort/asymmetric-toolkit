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

func main() {
	var spec cli.Specification = cli.Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument: map[string]cli.ArgumentDescriptor{
			"test": {
				1000, // >= 1000 is a project-defined FlagId
				cli.None,
				"noop default won't change",
				"This is an option.",
				cli.ParserNoop,
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
				cli.ParserInt(0,10),
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
				cli.ParserFloat(0.0,2.0),
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
				cli.ParserEnum("a","b","c"),
				cli.ExpectValue,
			},
		},
	}
	var ui cli.CommandLine
	exit, err := ui.Parse(&spec)
	if err != nil {
		fmt.Printf("Error:%v", err)
		os.Exit(cli.ErrArgumentParseError)
	}
	if exit {
		fmt.Printf("%v", err)
		os.Exit(cli.ErrArgumentParseError)
	}
}
