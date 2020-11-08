package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestCommandLine_Parse_Debug(t *testing.T) {
	defer func() { recover() }()
	args := []string{"--debug"}
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
				ParserFlag(),
				ExpectNone, //We expect None (no value) which will expect a flag in the end.
			},
		},
	}
	var ui CommandLine
	_, err := ui.Parse(&spec, args)
	if err != nil {
		panic(err)
	}

	errors.Assert(
		ui.Arguments[1003].String() == "myValue",
		fmt.Sprintf("Unexpected Value:%v", ui.Arguments[1000].String()))
}
