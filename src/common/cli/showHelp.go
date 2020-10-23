package cli

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/src/buildConfig"
	"fmt"
)

func (o *CommandLine) ShowHelp() {
	fmt.Printf("Starting %s (%s)...\n"+
		"\n"+
		"Copyright (c) %s %s.  All Rights Reserved.  <%s> \n"+
		"\n"+
		"Usage:\n",
		o.ProgramName,
		buildconfig.Version,
		buildconfig.CopyrightYear,
		buildconfig.Author,
		buildconfig.AuthorEmail)
	for k, v := range *o.spec {
		fmt.Printf(
			"\t%s\n"+
				"\t\t%s\n"+
				"\n",
			k,
			v.Help)
	}
}
