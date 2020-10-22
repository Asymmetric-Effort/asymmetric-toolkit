package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"fmt"
)

func ShowVersion() {
	fmt.Println(cli.Version)
}