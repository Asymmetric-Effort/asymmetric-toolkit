package cli

import "testing"

func TestCommandLine_Parse(t *testing.T) {
	args:=[]string{"flag","intVal","10","strVal","myValue"}
	var spec Specification = Specification{
		Author:      "Sam Caldwell",
		AuthorEmail: "mail@samcaldwell.net",
		ProgramName: "Asymmetric Effort",
		Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
		Version:     "0.0.1",
		Argument: map[string]ArgumentDescriptor{
			"flag": {
				1000, // >= 1000 is a project-defined FlagId
				None,
				"false",
				"This is a flag.",
				ParserBool(),
				ExpectFlag,
			},
			"intVal": {
				1001, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This is an integer value",
				ParserInt(),
				ExpectValue,
			},
			"intValDefault": {
				1002, // >= 1000 is a project-defined FlagId
				Integer,
				"5",
				"This tests default value (5)",
				ParserInt(0,10),
				ExpectValue,
			},
			"strVal": {
				1003, // >= 1000 is a project-defined FlagId
				String,
				"defaultValueShouldNotBeThere",
				"This is a string value.",
				ParserString(),
				ExpectValue,
			},
		},
	}
	var ui CommandLine
	_, err := ui.Parse(&spec,args)
	if err != nil {
		panic(err)
	}
}
