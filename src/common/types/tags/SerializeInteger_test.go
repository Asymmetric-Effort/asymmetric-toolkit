package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"fmt"
	"testing"
)

func TestTag_SerializeInteger(t *testing.T) {
	key := "name"
	value := 1
	o := NewInteger()
	o.Add(key, value)
	tag := o.SerializeInteger()
	expectedTag := []byte{1, 3, 4, 110, 97, 109, 101, 1, 0, 0, 0}
	fmt.Println("    :0.......8......F.")
	fmt.Println("tag1:", tag)
	fmt.Println("tag2:", expectedTag)
	errors.Assert(bytes.Equal(tag, expectedTag), fmt.Sprintf("Key-value mismatch. v:%s", string(tag)))
}
