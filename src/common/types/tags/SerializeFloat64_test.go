package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"
	"fmt"
	"testing"
)

func TestTag_SerializeFloat64(t *testing.T) {
	key := "name"
	value := 1.0
	o := NewFloat64()
	o.Add(key, value)
	tag := o.SerializeFloat64()
	expectedTag := []byte{1, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 240, 63}
	fmt.Println("    :0.......8......F.")
	fmt.Println("tag1:", tag)
	fmt.Println("tag2:", expectedTag)
	errors.Assert(bytes.Equal(tag, expectedTag), fmt.Sprintf("Key-value mismatch. v:%v", []byte(tag)))
}
