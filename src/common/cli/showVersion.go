package cli

import (
	buildconfig "asymmetric-effort/asymmetric-toolkit/src/buildConfig"
	"fmt"
)

func (o *Configuration) ShowVersion() {
	fmt.Println(buildconfig.Version)
}