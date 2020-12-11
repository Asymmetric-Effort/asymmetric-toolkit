package persistentTree

import (
	"os"
)

type Tree struct {
	header FileHeader
	file   *os.File
}
