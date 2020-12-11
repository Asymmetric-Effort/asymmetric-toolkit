package persistentTree

import (
	"fmt"
)

func (tree *Tree) gotoPosition(n int64) {
	_, err := tree.file.Seek(n, 0)
	if err != nil {
		panic(fmt.Sprintf("Tree::seek(): error moving to position %d [%x]:%v", n, n, err))
	}
}