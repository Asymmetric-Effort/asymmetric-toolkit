package encryption

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"bytes"
	"encoding/hex"
	"testing"
)

func TestCreateKey(t *testing.T){
	p:="6368616e676520746869732070617373"
	eKey, _ := hex.DecodeString(p)
	aKey:=*createKey(&p)
	errors.Assert(bytes.Equal(eKey,aKey),"Key mismatch")
}