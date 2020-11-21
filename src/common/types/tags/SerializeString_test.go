package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"bytes"

	//"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_SerializeString(t *testing.T) {
	for i := 0; i < 100; i++ {
		key := "name"
		value := "Value"
		o := NewString()
		o.Add(key, value)
		tag := o.SerializeString()
		expectedTag := []byte{1, 2, 4, 5, 110, 97, 109, 101, 86, 97, 108, 117, 101}
		fmt.Println("tag1:", tag)
		fmt.Println("tag2:", expectedTag)
		errors.Assert(bytes.Equal(tag, expectedTag), fmt.Sprintf("Key-value mismatch. v:%s", string(tag)))
	}
}
