package cli

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/src/buildConfig"
	"fmt"
)

func (o *CommandLine) ShowVersion() {
	fmt.Println(buildconfig.Version)
}