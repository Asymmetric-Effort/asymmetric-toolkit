package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestTree(t *testing.T) {
	var tree Tree
	errors.Assert(&tree != nil, "Expect successful instantiation only.")
}
