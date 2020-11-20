package tags

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
	"fmt"
)

func (o *String) Serialize() []byte {
	buf := bytes.Buffer{}
	for key, value := range *o {
		buf.Write(misc.Uint32ToByte(uint32(len(key))))   // Length of key
		buf.Write(misc.Uint32ToByte(uint32(len(value)))) // Length of value
		buf.Write([]byte(key))                           // key (bytes)
		buf.Write([]byte(value))                         // value (bytes)
	}
	fmt.Println("String::Serialize(): start d=", buf.Bytes())
	return buf.Bytes()
}
