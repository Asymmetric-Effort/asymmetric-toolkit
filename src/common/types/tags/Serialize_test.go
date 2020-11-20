package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestTag_StringSerialize(t *testing.T) {
	key := "name"
	value := "Value"
	o := NewString()
	o.Add(key, value)

	tag := o.Serialize()
	errors.Assert(string(tag) == fmt.Sprintf("%d%d%s:%s", len(key), len(value), key, value), fmt.Sprintf("Key-value mismatch. v:%s",string(tag)))
}
