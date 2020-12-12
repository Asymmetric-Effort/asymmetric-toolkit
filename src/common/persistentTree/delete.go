package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/file"
	"fmt"
	"os"
)

func deleteFile(n string) {
	if !file.FileExists(n) {
		panic(fmt.Sprintf("persistentTree::deleteFile() error deleting file (%s)", n))
	}
	err := os.Remove(n)
	if err != nil {
		panic(fmt.Sprintf("persistentTree::deleteFile() error cannot delete file (%s)", n))
	}
}
