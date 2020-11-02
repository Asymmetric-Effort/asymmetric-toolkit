package main

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
				"",
				"This is an option.",
				cli.ParserNoop,
				cli.ExpectFlag,
			},
			"myOption": {
				1001, // >= 1000 is a project-defined FlagId
				cli.None,
				"",
				"This is a Second Option",
				cli.ParserNoop,
				cli.ExpectValue,
			},
		},
	}
	var ui cli.CommandLine
	exit, err:=ui.Parse(&spec)
	if err!=nil {
		fmt.Printf("Error:%v",err)
		os.Exit(cli.ErrArgumentParseError)
	}
	if exit {
		fmt.Printf("%v",err)
		os.Exit(cli.ErrArgumentParseError)
	}
}
