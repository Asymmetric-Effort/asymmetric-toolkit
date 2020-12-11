package persistentTree

import (
	"os"
	"sync"
)

type Tree struct {
	mutex  sync.Mutex
	header FileHeader
	file   *os.File
}
