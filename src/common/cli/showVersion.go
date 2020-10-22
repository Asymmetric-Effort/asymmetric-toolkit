package cli

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/src/buildConfig"
	"fmt"
)

func ShowVersion() {
	fmt.Println(buildconfig.Version)
}